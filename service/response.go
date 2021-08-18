package service

import (
	"BullMock/parse"
	"encoding/json"
	"log"
	"net/http"
)

func ReturnData(mockInfo parse.MockInfo, req *http.Request, bodyBytes []byte) string {
	//那到请求头类型contentType
	var reqBody = new(interface{})
	if req.Header.Get("Content-Type") == parse.ApplicationJson {
		err := json.Unmarshal(bodyBytes, reqBody)
		if err != nil {
			return err.Error()
		}
		returnMap, ok := mockInfo.RespTemplate.(map[string]interface{})
		if ok {
			log.Println(returnMap)
		}
	}
	log.Println(reqBody)
	return ""
}
