package service

import (
	"BullMock/parse"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TAG = "|##|"
	ALL = "*"
)

func AddMock(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	mockInfo := &parse.MockInfo{}
	mockInstance := &parse.MockInstance{}
	err = json.Unmarshal(bodyBytes, mockInstance)
	if err != nil {
		fmt.Fprintf(w, "json错误%s", string(bodyBytes))
		return
	}
	mockInfo.Key = mockInstance.Url + TAG + mockInstance.Method
	if mockInstance.MockRule == "" {
		mockInfo.MockRule = ALL
	} else {
		mockInfo.MockRule = mockInstance.MockRule
	}
	mockInfo.Status = true
	mockInstance.AddMock(mockInfo)
}

func UpdateMock(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateMock")
}

func CloseMock(w http.ResponseWriter, r *http.Request) {
	log.Println("CloseMock")
}

func DeleteMock(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteMock")
}

func Accept(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	method := req.Method
	host := req.Host
	url := req.RequestURI
	key := host + url + TAG + method
	mockInfo := parse.MockCollect[key]

	if !mockInfo.Status {
		// todo 走真实请求 并返回
		return
	}

	if nil != mockInfo {
		w.Header().Set("Content-Type", mockInfo.RespContentType)
		w.Write([]byte(ReturnData(mockInfo.RespTemplate, req, bodyBytes)))
		return
	}
	w.Write([]byte("未在mock系统定义返回"))

}

func EveryHandler() http.Handler {
	return http.HandlerFunc(Accept)
}
