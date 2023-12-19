package transformer

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	python "github.pitagora/pkg/python3"
	"gitlab.com/technity/go-x/pkg/logger"
)

type Result struct {
	ID   string
	Data []byte
	Err  error
}

type Procedure struct {
	StepName      string              `json:"step-name"`
	ProcedureName string              `json:"procedure-name"`
	Entrypoint    string              `json:"entrypoint"`
	Arguments     map[string]Argument `json:"arguments"`
}

type Argument struct {
	Value interface{} `json:"value"`
	Type  interface{} `json:"type"`
}

func LoadPyFile(name string) (string, error) {
	return loadPyFile(name)
}

func loadPyFile(name string) (string, error) {
	// Open the file
	file, err := os.Open(filepath.Join("./python", name))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Read the file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	return string(data), nil
}

func Transform(ctx context.Context, proc Procedure, intermediateRes *sync.Map, results chan<- Result) {
	runtime.LockOSThread()

	fmt.Printf("Transform GIL %v\n", python.PyGILState_Check())
	fmt.Println(python.PyEval_ThreadsInitialized())
	fmt.Println(python.Py_IsInitialized())

	logger.GetLogger(ctx).WithFields(map[string]any{
		"type": "procedure",
		"name": proc.StepName,
	}).Info("transforming")

	startTime := time.Now()
	// Open the file
	source, err := loadPyFile(proc.ProcedureName)
	if err != nil {
		results <- Result{ID: proc.StepName, Err: err}
		return
	}
	fmt.Printf("source load '%s' took: %v\n", proc.StepName, time.Now().Sub(startTime))
	// state := python.PyGILState_Ensure()
	// defer python.PyGILState_Release(state)
	lock.Lock()
	defer lock.Unlock()

	startTime = time.Now()

	logger.GetLogger(ctx).Info("acquired")

	executor := NewPythonExecutor(proc.ProcedureName)
	defer executor.Cleanup()
	fmt.Printf("NewPythonExecutor load '%s' took: %v\n", proc.StepName, time.Now().Sub(startTime))
	startTime = time.Now()

	executor.CompileUserCode(source)
	fmt.Printf("CompileUserCode load '%s' took: %v\n", proc.StepName, time.Now().Sub(startTime))
	startTime = time.Now()
	res := executor.ExecuteUserCode(proc.Entrypoint, proc.Arguments, intermediateRes)
	fmt.Printf("ExecuteUserCode load '%s' took: %v\n", proc.StepName, time.Now().Sub(startTime))
	startTime = time.Now()

	// fmt.Println(string(res))
	// time.Sleep(200 * time.Millisecond)
	results <- Result{ID: proc.StepName, Data: res}
}
