package service

import (
	"log"
	"net/http"
)

type MockInstance struct {
	Method      string      `json:"method"`
	Url         string      `json:"url"`
	Request     interface{} `json:"request"`
	Response    interface{} `json:"response"`
	ContentType string      `json:"contentType"`
	ForwardUrl  string      `json:"forwardUrl"`
}

func AddMock(w http.ResponseWriter, r *http.Request) {
	originalHost := r.Host
	body := r.Body
	log.Println(originalHost, body)

}

func UpdateMock(w http.ResponseWriter, r *http.Request) {

}

func CloseMock(w http.ResponseWriter, r *http.Request) {

}

func DeleteMock(w http.ResponseWriter, r *http.Request) {

}

func Accept(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)
}
