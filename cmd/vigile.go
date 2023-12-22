package main

/*
#include "Python.h"
*/
import "C"
import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	python "github.pitagora/pkg/python3"
	"github.pitagora/pkg/transformer"
)

var Verbose, OutputJSON bool
var ctx = context.Background()

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Print verbose output")
	rootCmd.PersistentFlags().BoolVarP(&OutputJSON, "json", "j", false, "Print JSON output")

	err := godotenv.Load("/.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "vigile",
	Short: "Root command of vigile",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Initialize the Python interpreter
	runtime.LockOSThread()
	python.Py_Initialize()
	transformer.InitPythonExecutor()
	// _tstate := C.PyGILState_GetThisThreadState()
	// C.PyEval_ReleaseThread(_tstate)
	// defer python.Py_Finalize()

	l := sync.Mutex{}
	m := &sync.Map{}
	var wg sync.WaitGroup
	source, err := transformer.LoadPyFile("mocker.py")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		wg.Add(10)
		go func() {
			for j := 0; j < 10; j++ {
				name := fmt.Sprintf("%d-%d", i, j)
				go func() {
					defer wg.Done()
					// s := python.PyGILState_Ensure()
					// defer python.PyGILState_Release(s)
					l.Lock()
					defer l.Unlock()
					exec := transformer.NewPythonExecutor(name)
					exec.CompileUserCode(source)
					res := exec.ExecuteUserCode("Mocker", map[string]transformer.Argument{
						"data": {
							Value: "test",
						},
					}, m)
					fmt.Printf("result of %s: %s\n", name, string(res))
				}()
			}
		}()
	}
	wg.Wait()
	fmt.Println("done")

	Execute()
}

func printVerboseInput(srv, mthd string, data interface{}) {
	fmt.Println("Service:", srv)
	fmt.Println("Method:", mthd)
	fmt.Print("Input: ")
	printMessage(data)
}

func printMessage(data interface{}) {
	//var s string

	if OutputJSON {
	}

	fmt.Println(data)
}
