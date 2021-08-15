package main

import (
	"BullMock/service"
	"BullMock/service/pluginserver"
	"log"
	"net/http"
)

const (
	addMock    = "/mock/add"
	updateMock = "/mock/update"
	closeMock  = "/mock/close"
	deleteMock = "/mock/delete"
	invoke     = "/all"
)

func startMock() {
	r := pluginserver.NewRouter()
	//invokeRouter := r.PathPrefix(invoke).Subrouter()
	r.HandleFunc(addMock, service.AddMock).Methods("POST")
	r.HandleFunc(updateMock, service.UpdateMock).Methods("POST")
	r.HandleFunc(closeMock, service.CloseMock).Methods("GET")
	r.HandleFunc(deleteMock, service.DeleteMock).Methods("GET")
	r.HandleFunc(invoke, service.Accept)
	log.Fatal(http.ListenAndServe(":9090", r))
}

func main() {
	startMock()
	//log.Println(utils.Str(10))

	//compile := regexp.MustCompile("/invoke/[a-z0-9/\\?&=_A-Z]+")
	//matchString := compile.MatchString("/invoke/fff/ffff")
	//log.Println(matchString)

	var i interface{} = make(map[string]interface{})
	log.Printf("类型%T", i)
}
