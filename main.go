//usr/bin/go run $0 $@; exit $?

package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)


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
	objsModule := python.PyImport_ImportModule("objs")
	if objsModule == nil { panic("Error importing module: objs") }
	fmt.Printf("[MODULE REF] repr(objsModule) = %s\n", python.PyString_AS_STRING(objsModule.Repr()))

	obj := objsModule.GetAttrString("obj")
	if obj == nil { panic("[CLASS REF] Error importing object: obj") }
	fmt.Printf("[CLASS REF] repr(obj) = %s\n", python.PyString_AS_STRING(obj.Repr()))

	newObj := python.PyInstance_New(obj, python.PyTuple_New(0), nil)
	if newObj == nil { panic("[INSTANCE REF] Error instantiating object: obj") }
	fmt.Printf("[INSTANCE REF] repr(newObj) = %s\n", python.PyString_AS_STRING(newObj.Repr()))

	methodCall := newObj.CallMethod("obj_method")
	if newObj == nil { panic("[METHOD CALL REF] Error calling object method: obj_method") }
	fmt.Printf("[METHOD CALL REF] repr(methodCall) return value = %s\n", python.PyString_AS_STRING(methodCall.Repr()))
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