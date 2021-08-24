package main

import (
	"BullMock/utils"
	"log"
	"reflect"
	"strings"
)

func main() {
	//var data interface{} = map[string]string{
	//	"a": "b",
	//}
	//log.Println(test())v
	var ii = 10
	log.Println(InvokeFun("str", ii))
}

func test() interface{} {
	refStruct := &utils.Func{}
	refSS := reflect.ValueOf(refStruct)
	i := refSS.MethodByName("Str")
	i2 := i.Call([]reflect.Value{reflect.ValueOf(10)})[0].Interface()
	log.Println(i2)
	return i2
}

func InvokeFun(funcName string, params ...interface{}) interface{} {
	var resValue reflect.Value
	f := &utils.Func{}
	resValue = reflect.ValueOf(f)
	invokeName := strings.ToUpper(funcName[0:1]) + funcName[1:]
	function := resValue.MethodByName(invokeName)
	var refParam []reflect.Value
	if len(params) == 1 && params[0] == nil {
		return function.Call(nil)[0]
	} else {
		refParam = make([]reflect.Value, len(params))
		for i := 0; i < len(params); i++ {
			refParam[i] = reflect.ValueOf(params[i])
		}
		log.Println(function.Call(refParam))
		return "values[0].Interface()"
	}
}

//func test1() {
//	// 1. 数据准备把字符串转换到对象内存储
//	var json_data interface{}
//	json.Unmarshal([]byte(dataStr), &json_data)
//
//	res, err := jsonpath.JsonPathLookup(json_data, "$.expensive")
//	if err == nil {
//		fmt.Println("step 1 res: $.expensive")
//		fmt.Println(res)
//	}
//}
//
//func test2() {
//	var json_data interface{}
//	json.Unmarshal([]byte(dataStr), &json_data)
//	//or reuse lookup pattern
//	pat, _ := jsonpath.Compile(`$.store.book[?(@.price < $.expensive)].price`)
//	res, _ := pat.Lookup(json_data)
//	fmt.Println("step 2 res:")
//	fmt.Println(res)
//}
//
//func test3() {
//	// 3. 未找到对象
//	var json_data interface{}
//	json.Unmarshal([]byte(dataStr), &json_data)
//
//	res, err := jsonpath.JsonPathLookup(json_data, "$.expensive1")
//	if err == nil {
//		fmt.Println("step 3 res: $.expensive")
//		fmt.Println(res)
//	} else {
//		fmt.Printf("dddd: %v ", err)
//	}
//}
//
//var dataStr string = `
//{
//    "store": {
//        "book": [
//            {
//                "category": "reference",
//                "author": "Nigel Rees",
//                "title": "Sayings of the Century",
//                "price": 8.95
//            },
//            {
//                "category": "fiction",
//                "author": "Evelyn Waugh",
//                "title": "Sword of Honour",
//                "price": 12.99
//            },
//            {
//                "category": "fiction",
//                "author": "Herman Melville",
//                "title": "Moby Dick",
//                "isbn": "0-553-21311-3",
//                "price": 8.99
//            },
//            {
//                "category": "fiction",
//                "author": "J. R. R. Tolkien",
//                "title": "The Lord of the Rings",
//                "isbn": "0-395-19395-8",
//                "price": 22.99
//            }
//        ],
//        "bicycle": {
//            "color": "red",
//            "price": 19.95
//        }
//    },
//    "expensive": 10
//}
//`
