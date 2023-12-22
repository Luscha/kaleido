package transformer

import (
	"fmt"
	"math"
	"reflect"
	"sync"

	"github.pitagora/pkg/node"
	python "github.pitagora/pkg/python3"
)

// PythonExecutor is a struct representing a Python executor with a restricted environment.
type PythonExecutor struct {
	name         string
	compiledCode *python.PyObject
	objects      []*python.PyObject
}

var globals *python.PyObject
var compileRestricted *python.PyObject
var lock sync.Mutex

func InitPythonExecutor() (*python.PyObject, *python.PyObject) {
	// globals
	source, err := loadPyFile("setup.py")
	if err != nil {
		panic(err)
	}
	// TODO catch errors
	// Run setupCode to establish the restricted environment
	python.PyRun_SimpleString(source)

	// Fetch the safe_globals dictionary
	mainModule := python.PyImport_AddModule("__main__")
	if mainModule == nil {
		python.PyErr_Print()
		panic("Failed to get __main__ module")
	}
	mainDict := python.PyModule_GetDict(mainModule)
	globals = python.PyDict_GetItemString(mainDict, "safe_globals")
	if globals == nil {
		python.PyErr_Print()
		panic("Failed to get safe_globals")
	}
	globals.IncRef()

	// restricted compile
	// Load RestrictedPython module
	restrictedPythonModule := python.PyImport_ImportModule("RestrictedPython")
	if restrictedPythonModule == nil {
		panic("Failed to import RestrictedPython")
	}
	// defer restrictedPythonModule.DecRef()

	// Get the compile_restricted function
	compileRestricted = restrictedPythonModule.GetAttrString("compile_restricted")
	if compileRestricted == nil {
		panic("Failed to get compile_restricted")
	}
	compileRestricted.IncRef()
	// defer compileRestricted.DecRef()
	return globals, compileRestricted
}

// NewPythonExecutor initializes and returns a new PythonExecutor.
func NewPythonExecutor(name string) *PythonExecutor {
	// InitPythonExecutor()
	executor := &PythonExecutor{
		name:    name,
		objects: make([]*python.PyObject, 0),
	}

	return executor
}

func (executor *PythonExecutor) AddObject(obj *python.PyObject) {
	executor.objects = append(executor.objects, obj)
}

func (executor *PythonExecutor) Cleanup() {
	// ! for some reason trying to decrefs on objects instantiated within the executor
	// ! causes random SIGSEV
	// ! investigate further
	// for _, o := range executor.objects {
	// 	o.DecRef()
	// }
	// executor.compiledCode.DecRef()
}

// CompileUserCode compiles the given user code in the restricted environment.
func (executor *PythonExecutor) CompileUserCode(userCode string) {
	fmt.Printf("GIL: %v\n", python.PyGILState_Check())
	fmt.Printf("-----CompileUserCode : %s\n", executor.name)
	// Define and compile userCode
	pyUserCode := python.PyUnicode_FromString(userCode)
	executor.AddObject(pyUserCode)

	pyFileName := python.PyUnicode_FromString("<string>")
	executor.AddObject(pyFileName)

	args := python.PyTuple_New(2)
	python.PyTuple_SetItem(args, 0, pyUserCode)
	python.PyTuple_SetItem(args, 1, pyFileName)

	fmt.Println("compileRestricted:", python.PyUnicode_AsUTF8(python.PyObject_Repr(compileRestricted)))
	// fmt.Println("compileRestricted:", python.PyUnicode_AsUTF8(python.PyObject_Repr(args)))
	executor.compiledCode = compileRestricted.Call(args, nil)
	if executor.compiledCode == nil {
		panic("Failed to compile user code")
	}
}

// ExecuteUserCode executes the given compiled user code.
func (executor *PythonExecutor) ExecuteUserCode(entrypoint string, args map[string]Argument, intermediateRes *sync.Map) []byte {
	if executor.name == "linear-regression.py" {
		fmt.Println("here")
	}
	fmt.Printf("-----ExecuteUserCode : %s\n", executor.name)
	pyargs := executor.BuildUserCodeArgs(args, intermediateRes)

	localEnv := python.PyDict_New()
	executor.AddObject(localEnv)

	// ! TODO investigate how to make import work
	// if the script imports something like
	//     from sklearn.linear_model import LinearRegression as lr
	// outside the function, the import gets correctly loaded in localEnv but is not availbale in the function
	//
	// from sklearn.linear_model import LinearRegression as lr
	// def LinearRegression(params):
	//		model = lr()
	//
	// ^ will throw
	//
	// def LinearRegression(params):
	// 		from sklearn.linear_model import LinearRegression as lr
	//		model = lr()
	// ^ won't throw
	// Execute the compiled user code
	if python.PyEval_EvalCode(executor.compiledCode, globals, localEnv) == nil {
		fmt.Println("Error running the compiled code")
		python.PyErr_Print()
		return nil
	}

	// Retrieve the userFunction function from localEnv
	fn := python.PyDict_GetItemString(localEnv, entrypoint)
	if fn == nil {
		fmt.Println("Function userFunction not found")
		return nil
	}
	executor.AddObject(fn)

	fmt.Println("fn:", python.PyUnicode_AsUTF8(python.PyObject_Repr(fn)))
	// fmt.Println("pyargs:", python.PyUnicode_AsUTF8(python.PyObject_Repr(pyargs)))
	fmt.Println("localEnv:", python.PyUnicode_AsUTF8(python.PyObject_Repr(localEnv)))

	result := fn.CallFunctionObjArgs(pyargs)
	if python.PyErr_Occurred() != nil {
		fmt.Printf("Error calling the %s function\n", entrypoint)
		ptype, pvalue, ptraceback := python.PyErr_Fetch()
		fmt.Println("ptype:", python.PyUnicode_AsUTF8(python.PyObject_Repr(ptype)))
		fmt.Println("pvalue:", python.PyUnicode_AsUTF8(python.PyObject_Repr(pvalue)))
		if ptraceback != nil {
			// TODO
			// Import the traceback module
			// traceback := C.PyImport_ImportModule(C.CString("traceback"))
			// format_exception := C.PyObject_GetAttrString(traceback, C.CString("format_exception"))

			// // Get the formatted stack trace
			// args := C.PyTuple_New(3)
			// C.PyTuple_SetItem(args, 0, ptype)
			// C.PyTuple_SetItem(args, 1, pvalue)
			// C.PyTuple_SetItem(args, 2, ptraceback)
			// pyStacktrace := C.PyObject_CallObject(format_exception, args)
			// stacktraceStr := C.GoString(C.PyUnicode_AsUTF8(C.PyUnicode_Join(C.PyUnicode_FromString(C.CString("\n")), pyStacktrace)))

			// // Print the stack trace in Go
			// fmt.Println("Python Stack Trace:", stacktraceStr)
		}
		executor.AddObject(ptype)
		executor.AddObject(pvalue)
		executor.AddObject(ptraceback)
		python.PyErr_Print()
		return nil
	}
	executor.AddObject(result)

	// Handle the result
	if python.PyUnicode_Check(result) {
		resultStr := python.PyUnicode_AsUTF8(result)
		fmt.Println("Result:", len(resultStr))
		return []byte(resultStr)
	} else {
		fmt.Println("Result is not a string")
	}
	return nil
}

func (executor *PythonExecutor) BuildUserCodeArgs(args map[string]Argument, intermediateRes *sync.Map) *python.PyObject {
	// Create a new Python dictionary
	pyArgs := python.PyDict_New()
	executor.AddObject(pyArgs)

	for key, arg := range args {
		pyValue := executor.convertToPyObjectWrapper(arg.Value, intermediateRes)

		// Set the Python object in the pyArgs dictionary
		python.PyDict_SetItemString(pyArgs, key, pyValue)
	}

	return pyArgs
}

func (executor *PythonExecutor) convertToPyObjectWrapper(value interface{}, intermediateRes *sync.Map) *python.PyObject {
	pyItem, singleton := executor.convertToPyObject(value, intermediateRes)
	if !singleton {
		executor.AddObject(pyItem)
	}
	return pyItem
}

func (executor *PythonExecutor) convertToPyObject(value interface{}, intermediateRes *sync.Map) (*python.PyObject, bool) {
	// Convert value to a PyObject based on inferred type
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Slice, reflect.Array:
		// Handle slices and arrays
		s := reflect.ValueOf(value)
		pyList := python.PyList_New(s.Len())
		for i := 0; i < s.Len(); i++ {
			item := s.Index(i).Interface()
			pyItem := executor.convertToPyObjectWrapper(item, intermediateRes)
			python.PyList_SetItem(pyList, i, pyItem)
		}
		return pyList, false

	case reflect.String:
		// Handle strings, including 'data' key special case
		strValue := value.(string)
		specialName := node.TypeAndStringKey(node.GetNameAndType(strValue))
		if data, exists := intermediateRes.Load(specialName); exists {
			return python.PyUnicode_FromString(string(data.([]byte))), false
		} else {
			return python.PyUnicode_FromString(strValue), false
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Handle integers
		return python.PyLong_FromLongLong(reflect.ValueOf(value).Int()), false

	case reflect.Float32, reflect.Float64:
		// Handle floats
		floatValue := reflect.ValueOf(value).Float()

		// Check if the float value is a whole number and within the int64 range
		if floatValue == math.Trunc(floatValue) && floatValue <= math.MaxInt64 && floatValue >= math.MinInt64 {
			// Convert to int64 and create a Python integer object
			return python.PyLong_FromLongLong(int64(floatValue)), false
		} else {
			// Create a Python float object
			return python.PyFloat_FromDouble(floatValue), false
		}

	case reflect.Bool:
		// Handle booleans
		boolValue := value.(bool)
		if boolValue {
			return python.Py_True, true
		} else {
			return python.Py_False, true
		}

	case reflect.Map:
		// Handle maps (objects)
		mapValue := reflect.ValueOf(value)
		pyDict := python.PyDict_New()
		for _, key := range mapValue.MapKeys() {
			keyStr, ok := key.Interface().(string)
			if !ok {
				fmt.Printf("Non-string key in map: %v\n", key)
				continue
			}
			val := mapValue.MapIndex(key).Interface()
			pyItem := executor.convertToPyObjectWrapper(val, intermediateRes)
			python.PyDict_SetItemString(pyDict, keyStr, pyItem)
		}
		return pyDict, false

	default:
		// If the type is not recognized, return None
		fmt.Printf("Unsupported type: %v\n", valueType)
		return python.Py_None, true
	}
}
