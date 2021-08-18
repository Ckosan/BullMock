package main

import (
	"BullMock/utils"
	"log"
	"reflect"
)

func main() {
	//var data interface{} = map[string]string{
	//	"a": "b",
	//}
	refStruct := &utils.Func{}
	refSS := reflect.ValueOf(refStruct)
	log.Println(refSS.NumMethod())
	log.Println(refSS.Method(1))
	i := refSS.MethodByName("Str")
	log.Println(i.Call([]reflect.Value{reflect.ValueOf(10)}))
	log.Println()

}
