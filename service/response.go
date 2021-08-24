package service

import (
	"BullMock/parse"
	"encoding/json"
	"github.com/oliveagle/jsonpath"
	"net/http"
	"strings"
)

func ReturnData(mockInfo parse.MockInfo, req *http.Request, bodyBytes []byte) ([]byte, error) {
	//那到请求头类型contentType
	var jsonReqBody = make(map[string]interface{}, 16)
	if req.Header.Get("Content-Type") == parse.ApplicationJson {
		err := json.Unmarshal(bodyBytes, &jsonReqBody)
		if err != nil {
			return nil, err
		}
		returnMap, ok := mockInfo.RespTemplate.(map[string]interface{})
		copyMap := make(map[string]interface{}, 16)
		for k, v := range returnMap {
			copyMap[k] = v
		}
		if ok {
			for k, v := range copyMap {
				exprFuncValue, isExpr := parse.FindFunctionExpression(v)
				if isExpr {
					exprValueStr := exprFuncValue.(string)
					funcName := exprValueStr[0:strings.Index(exprValueStr, "(")]
					copyMap[k] = parse.InvokeFun(funcName, parse.FindFuncParam(exprValueStr))
				}
				exprValue, isExpr := parse.FindValueExpression(v)
				if isExpr {
					exprValueStr := exprValue.(string)
					value, err := replaceValue(exprValueStr, &jsonReqBody, &returnMap)
					if err == nil {
						copyMap[k] = value
					} else {
						return nil, err
					}
				}

			}
			data, _ := json.Marshal(copyMap)
			return data, nil
		}
	}
	return nil, nil
}

func SyncCreateResponseData(mockInfo parse.MockInfo) {

}

func replaceValue(s string, reqBody *map[string]interface{}, responseMap *map[string]interface{}) (interface{}, error) {
	if strings.ContainsAny(s, "fromReq") {
		s = strings.ReplaceAll(s, "fromReq", "$")
		value, err := jsonpath.JsonPathLookup(*reqBody, s)
		if err == nil {
			return value, nil
		} else {
			return nil, err
		}
	} else {
		s = "$." + s
		value, err := jsonpath.JsonPathLookup(*responseMap, s)
		if err == nil {
			return value, nil
		} else {
			return nil, err
		}
	}
}
