package service

import (
	"BullMock/parse"
	"BullMock/utils"
	"encoding/json"
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
	w.Header().Set("Content-Type", parse.ApplicationJson)
	if err != nil {
		errData := Data{
			Msg:  err.Error(),
			Code: 2001,
		}
		dataByte, _ := json.Marshal(errData)
		w.Write(dataByte)
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
	returnData := Data{
		Msg:  "success",
		Code: 2000,
	}
	dataByte, _ := json.Marshal(returnData)
	w.Write(dataByte)
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
	key := scheme(req) + host + url + TAG + method
	var mockInfo = parse.MockCollect[key]

	if mockInfo == nil {
		// 走真实请求 并返回
		err, resp := utils.DoHttp(req)
		if err == nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			for k, v := range resp.Header {
				w.Header().Set(k, v[0])
			}
			w.Write(bytes)
		} else {
			w.Write([]byte(err.Error()))
		}
		return
	}

	if nil != mockInfo && mockInfo.Status {
		w.Header().Set("Content-Type", mockInfo.RespContentType)
		data, err := ReturnData(*mockInfo, req, bodyBytes)
		if err == nil {
			w.Write(data)
		} else {
			orgResp, _ := json.Marshal(mockInfo.RespTemplate)
			w.Write(orgResp)
		}
		return
	}
	returnData := Data{
		Msg:  "未在mock系统定义返回或开关未打开",
		Code: 4002,
	}
	marshal, _ := json.Marshal(returnData)
	w.Header().Set("Content-Type", parse.ApplicationJson)
	w.Write(marshal)
}

type Data struct {
	Msg  string                 `json:"msg"`
	Code uint                   `json:"code"`
	Data map[string]interface{} `json:"data"`
}

func scheme(req *http.Request) string {
	if req.TLS == nil {
		return utils.HTTP
	}
	return utils.HTTPS
}
