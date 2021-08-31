package service

import (
	"BullMock/models"
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
	log.Printf("增加mock:%s成功", string(bodyBytes))
	go func() {
		_, err := models.GetMockIns(mockInstance.Method, mockInstance.Url)
		if err == nil {
			return
		}
		err = models.CreateMockIns(mockInstance)
		if err != nil {
			return
		}
	}()
}

func UpdateMock(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	mockInstance := &parse.MockInstance{}
	err = json.Unmarshal(bodyBytes, mockInstance)
	models.UpdateMockIns(mockInstance)
	log.Println("UpdateMock")
}

func CloseMock(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	host := req.Host
	url := req.RequestURI
	key := scheme(req) + host + url + TAG + method
	var mockInfo = parse.MockCollect[key]
	mockInfo.Status = false
	models.UpdateMockStatus(false, method, url)
	log.Println("CloseMock")
}

func DeleteMock(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteMock")
}

func AddScript(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	script := models.Script{}
	err = json.Unmarshal(bodyBytes, &script)
	if err != nil {
		log.Println(err)
		return
	}
	models.AddScript(&script)
}

func UpScript(w http.ResponseWriter, r *http.Request) {

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

	if mockInfo.Status {
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
