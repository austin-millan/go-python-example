//usr/bin/go run $0 $@; exit $?

package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

var PyStr = python.PyString_FromString


func import_and_call_func() {
	funcsModule := python.PyImport_ImportModule("funcs")
	if funcsModule == nil { panic("[MODULE REF] Error importing module: funcs.py") }
	fmt.Printf("[MODULE REF] repr(funcsModule) = %s\n", python.PyString_AS_STRING(funcsModule.Repr()))

	funcName := funcsModule.GetAttrString("foo")
	if funcName == nil { panic("[FUNCTION REF] Error importing function: foo") }
	fmt.Printf("[FUNCTION REF] repr(funcName) = %s\n", python.PyString_AS_STRING(funcName.Repr()))

	funcCall := funcName.Call(python.PyTuple_New(0), python.PyDict_New())
	if funcCall == nil { panic("[FUNCTION CALL REF] Error calling function: foo") }
	fmt.Printf("[FUNCTION CALL REF] repr(methodCall) = %s\n", python.PyString_AS_STRING(funcCall.Repr()))
}

func import_and_use_obj() {
	// Import class from module
	objsModule := python.PyImport_ImportModule("objs")
	if objsModule == nil { panic("Error importing module: objs") }
	fmt.Printf("[MODULE REF] repr(objsModule) = %s\n", python.PyString_AS_STRING(objsModule.Repr()))

	obj := objsModule.GetAttrString("obj")
	if obj == nil { panic("[CLASS REF] Error importing object: obj") }
	fmt.Printf("[CLASS REF] repr(obj) = %s\n", python.PyString_AS_STRING(obj.Repr()))

	// Instantiate obj
	newObj := python.PyInstance_New(obj, python.PyTuple_New(0), nil)
	if newObj == nil { panic("[INSTANCE REF] Error instantiating object: obj") }
	fmt.Printf("[INSTANCE REF] repr(newObj) = %s\n", python.PyString_AS_STRING(newObj.Repr()))

	// Call method (no arguments)
	methodCall := newObj.CallMethod("obj_method")
	if newObj == nil { panic("[METHOD CALL REF] Error calling object method: obj_method") }
	fmt.Printf("[METHOD CALL REF] repr(methodCall) return value = %s\n", python.PyString_AS_STRING(methodCall.Repr()))

	// Create Python tuple and call method with tuple as arg
	methodArg := python.PyTuple_New(1)
	err := python.PyTuple_SetItem(methodArg, 0,  PyStr("FooArg"))
	if err != nil { panic("[NEW TUPLE REF] Error setting item in tuple: methodArg") }
	methodCallArg := newObj.CallMethodObjArgs("obj_method_arg", methodArg)
	if methodCallArg == nil { panic("[METHOD CALL REF] Error calling object method: obj_method_arg") }
	fmt.Printf("[METHOD CALL REF] repr(methodCallArg) return value = %s\n", python.PyString_AS_STRING(methodCallArg.Repr()))

	// Now try with multiple args
	methodArgTwo := python.PyList_New(2)
	err = python.PyList_SetItem(methodArgTwo, 0,  PyStr("ListElement0"))
	if err != nil { panic("[LIST SET REF] Error setting element at list[0]: methodArgTwo") }
	err = python.PyList_SetItem(methodArgTwo, 1,  PyStr("ListElement1"))
	if err != nil { panic("[LIST SET REF] Error setting element at list[1]: methodArgTwo") }
	methodCallArgs := newObj.CallMethodObjArgs("obj_method_args", methodArg, methodArgTwo)
	if methodCallArgs == nil { panic("[METHOD CALL REF] Error calling object method: obj_method_args") }
	fmt.Printf("[METHOD CALL REF] repr(methodCallArgs) return value = %s\n", python.PyString_AS_STRING(methodCallArgs.Repr()))
}


func main() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
	defer python.Finalize()
	import_and_call_func()
	import_and_use_obj()
	python.Finalize()
}