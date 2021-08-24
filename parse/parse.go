package parse

import (
	"BullMock/utils"
	"container/list"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	ApplicationJson       = "application/json"
	ApplicationUrlencoded = "application/x-www-form-urlencoded"
	MultipartFormData     = "multipart/form-data"
	TextPlain             = "text/plain"
	ApplicationXml        = "application/xml"
	regMatch              = "\\{\\{.*\\}\\}"
	regFunc               = "\\$.*\\)"
)

type Parse interface {
	ParseBody() error
	ParseUrl() error
}

type MockInstance struct {
	Method          string      `json:"method"`
	Url             string      `json:"url"`
	Request         interface{} `json:"request"`
	Response        interface{} `json:"response"`
	ReqContentType  string      `json:"reqContentType"`
	RespContentType string      `json:"respContentType"`
	ForwardUrl      string      `json:"forwardUrl"`
	MockRule        string      `json:"mockRule"`
}

type MockInfo struct {
	Status          bool
	Key             string
	RespTemplate    interface{}
	MockRule        string
	RespContentType string
	Expressions     []*Expression
}

type Expression struct {
	valueExp, funcExp, ruleExp string
}

var MockCollect = make(map[string]*MockInfo, 16)

func (instance *MockInstance) ParseBody(req *http.Request) error {

	return nil
}

func (instance *MockInstance) AddMock(mocInfo *MockInfo) error {
	mocInfo.RespContentType = instance.RespContentType
	//if instance.ReqContentType == ApplicationJson {
	//	request := instance.Request
	//	reqMap, ok := request.(map[string]interface{}) // interface{}转map
	//	if ok {
	//		for k, v := range reqMap {
	//			log.Println(k, v)
	//		}
	//	}
	//
	//}
	if instance.RespContentType == ApplicationJson {
		respMap, ok := instance.Response.(map[string]interface{})
		var exps = make([]*Expression, 16, 16)
		if ok {
			//i := 0
			//for _, v := range respMap {
			//	exp := &Expression{}
			//FindValueExpression(v, exp)
			//FindFunctionExpression(v, exp)
			//exps[i] = exp
			//i++
			mocInfo.RespTemplate = respMap
		}
		mocInfo.Expressions = exps
	}
	MockCollect[mocInfo.Key] = mocInfo
	if strings.HasPrefix(instance.ReqContentType, MultipartFormData) {

	}

	if instance.ReqContentType == ApplicationUrlencoded {
	}
	return nil
}

func (instance *MockInstance) ParseUrl() error {
	return nil
}

func FindValueExpression(val interface{}) (interface{}, bool) {
	if reflect.TypeOf(val).Name() == "string" {
		valStr := val.(string)
		reg := regexp.MustCompile(regMatch)
		if reg.MatchString(valStr) {
			ret := reg.Find([]byte(valStr))
			express := ret[2 : len(ret)-2]
			//exp.valueExp = string(express[:])
			return string(express[:]), true
		}
	}
	return val, false
}

func FindFunctionExpression(val interface{}) (interface{}, bool) {
	if reflect.TypeOf(val).Name() == "string" {
		valStr := val.(string)
		reg := regexp.MustCompile(regFunc)
		if reg.MatchString(valStr) {
			ret := reg.Find([]byte(valStr))
			express := ret[1:]
			//exp.funcExp = string(express[:])
			return string(express[:]), true
		}
	}
	return val, false
}

func FindFuncParam(str string) *list.List {
	indexByte := strings.IndexByte(str, '(')
	param := str[indexByte+1 : len(str)-1]
	if param == "" {
		return nil
	}
	splits := strings.Split(param, ",")
	params := list.New()
	for i := 0; i < len(splits); i++ {
		if strings.HasPrefix(splits[i], "'") && strings.HasSuffix(splits[i], "'") {
			params.PushBack(splits[i][1 : len(splits[i])-1])
		} else {
			parseInt, err := strconv.ParseInt(splits[i], 10, 8)
			if err != nil {
				parseValue, errBool := strconv.ParseBool(splits[i])
				if errBool != nil {
					parseF, errF := strconv.ParseFloat(splits[i], 32)
					if errF != nil {

					} else {
						params.PushBack(parseF)
					}
				} else {
					params.PushBack(parseValue)
				}
			} else {
				params.PushBack(int(parseInt))
			}

		}
	}
	return params
}

//	FindRuleValue 使用jsonPath 格式 如 $.request.name=='Pierson' json字段匹配
// 	$.fromKey.name== 'Pierson'   表单name字段为Pierson
//	$.header.Cookie=='uu=svn' 	 header里面Cookie的匹配
//	$.urlPath=='/v1/get-info'    urlPath匹配
func FindRuleValue(s string) {

}

var FuncCollect = make(map[string]interface{}, 16)

func init() {

}

func InvokeFun(funcName string, param *list.List) interface{} {
	var resValue reflect.Value
	f := &utils.Func{}
	resValue = reflect.ValueOf(f)
	invokeName := strings.ToUpper(funcName[0:1]) + funcName[1:]
	function := resValue.MethodByName(invokeName)
	var refParam []reflect.Value
	if param == nil {
		return function.Call(nil)[0]
	} else {
		refParam = make([]reflect.Value, param.Len())
		for i := 0; i < param.Len(); i++ {
			refParam[i] = reflect.ValueOf(param.Front().Value)
		}
		log.Println()
		return function.Call(refParam)[0].Interface()
	}
}

//func transfer(v interface{}) {
//	switch t := v.(type) {
//	case string:
//	case int:
//	case float32:
//	case float64:
//
//	}
//}
