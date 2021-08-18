package parse

import (
	"log"
	"net/http"
	"reflect"
	"regexp"
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
	if instance.ReqContentType == ApplicationJson {
		request := instance.Request
		reqMap, ok := request.(map[string]interface{}) // interface{}转map
		if ok {
			for k, v := range reqMap {
				log.Println(k, v)
			}
		}

	}
	if instance.RespContentType == ApplicationJson {
		respMap, ok := instance.Response.(map[string]interface{})
		var exps = make([]*Expression, 16, 16)
		if ok {
			i := 0
			for _, v := range respMap {
				exp := &Expression{}
				findValueExpression(v, exp)
				findFunctionExpression(v, exp)
				exps[i] = exp
				i++
			}
			mocInfo.RespTemplate = respMap
			mocInfo.Expressions = exps
		}

	}
	if strings.HasPrefix(instance.ReqContentType, MultipartFormData) {

	}

	if instance.ReqContentType == ApplicationUrlencoded {
	}
	return nil
}

func (instance *MockInstance) ParseUrl() error {
	return nil
}

func findValueExpression(val interface{}, exp *Expression) interface{} {
	if reflect.TypeOf(val).Name() == "string" {
		valStr := val.(string)
		reg := regexp.MustCompile(regMatch)
		if reg.MatchString(valStr) {
			ret := reg.Find([]byte(valStr))
			express := ret[2 : len(ret)-2]
			exp.valueExp = string(express[:])
			return string(express[:])
		}
	}
	return val
}

func findFunctionExpression(val interface{}, exp *Expression) interface{} {
	if reflect.TypeOf(val).Name() == "string" {
		valStr := val.(string)
		reg := regexp.MustCompile(regFunc)
		if reg.MatchString(valStr) {
			ret := reg.Find([]byte(valStr))
			express := ret[1 : len(ret)-2]
			exp.funcExp = string(express[:])
			return string(express[:])
		}
	}
	return val
}

//	findRuleValue 使用jsonPath 格式 如 $.request.name=='Pierson' json字段匹配
// 	$.fromKey.name== 'Pierson'   表单name字段为Pierson
//	$.header.Cookie=='uu=svn' 	 header里面Cookie的匹配
//	$.urlPath=='/v1/get-info'    urlPath匹配
func findRuleValue(s string) {

}

var FuncCollect = make(map[string]interface{}, 16)

func loadFunc() {

}
