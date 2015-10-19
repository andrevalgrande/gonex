package exceptions

import (
	"errors"
	"log"
	"runtime"
	"time"
)
import "fmt"

//gone relies on a singleton object. Evidently, there's room for improvement here
type singleton struct {
	Stack    string
	Errors   []string
	Warnings []string
}

var instance *singleton = nil
var panicKey = "f1c7501f-1942-42c3-a302-4856f1d997b0"

func resetInstance() {
	instance = new(singleton)
}

func getInstance() *singleton {
	if instance == nil {
		resetInstance()
	}
	return instance
}

//appends a textual warning into the exception
func AppendWarning(warning string) {
	getInstance()
	var contents = fmt.Sprintf("Aviso - %v: %v\n", time.Now(), warning)
	instance.Warnings = append(instance.Warnings, contents)
}

//appends a native go error into the exception
func AppendError(err error) {
	getInstance()
	var contents = fmt.Sprintf("Erro - %v: %v\n", time.Now(), err)
	instance.Errors = append(instance.Errors, contents)
}

//throws an exception with a message. The message will be made into a native go error
func Throw(message string) {
	getInstance()
	AppendError(errors.New(message))
	instance.Stack = getStack()
	log.Println(fmt.Sprintf("%v Exception raised: ", time.Now()))
	log.Println(GetString(true))
	panic(panicKey)
}

func getStack() string {
	stack := make([]byte, 1000)
	n := runtime.Stack(stack, false)
	stack = make([]byte, n)
	runtime.Stack(stack, false)
	return string(stack)
}

//catches a gone exception and runs the code within the fn parameter
//MUST be called as a defer
func Catch(fn func()) {
	r := recover()
	if r != nil {
		if r == panicKey {
			fn()
			Clear()
		} else {
			panic(r)
		}
	}
}

//gets the warnings and errors in the gone exception as a (poorly) formatted string
//the "includeStack" parameter allows the method to also return the callstack at the moment of the throw
func GetString(includeStack bool) (output string) {
	output = ""
	for _, warning := range getInstance().Warnings {
		output += warning
	}

	for _, error := range getInstance().Errors {
		output += error
	}

	if includeStack {
		output += getInstance().Stack
	}

	return
}

//gets the warnings and errors in the gone exception as a native error slice
//the "includeStack" parameter allows the method to also return the callstack at the moment of the throw
func GetErrorSlice(includeStack bool) (output []error) {
	getInstance()
	for _, warning := range instance.Warnings {
		output = append(output, errors.New(warning))
	}

	for _, err := range instance.Errors {
		output = append(output, errors.New(err))
	}

	if includeStack {
		output = append(output, errors.New(instance.Stack))
	}
	return
}

//resets the singleton instance, clearing errors and warnings
func Clear() {
	resetInstance()
}
