package main

import (
	"BullMock/service"
	"BullMock/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	mock       = "/mock"
	addMock    = "/add"
	updateMock = "/update"
	closeMock  = "/close"
	deleteMock = "/delete"
	invoke     = "/{[a-z0-9/\\?&=_A-Z]+}"
	//invoke     = "/invoke"
)

func startMock() {
	r := mux.NewRouter()
	subRouter := r.PathPrefix(mock).Subrouter()
	//invokeRouter := r.PathPrefix(invoke).Subrouter()
	subRouter.HandleFunc(addMock, service.AddMock).Methods("POST")
	subRouter.HandleFunc(updateMock, service.UpdateMock).Methods("POST")
	subRouter.HandleFunc(closeMock, service.CloseMock).Methods("GET")
	subRouter.HandleFunc(deleteMock, service.DeleteMock).Methods("GET")
	r.HandleFunc(invoke, service.Accept)
	log.Fatal(http.ListenAndServe(":9090", r))
}

func main() {
	startMock()
	log.Println(utils.Str(10))

	//compile := regexp.MustCompile("/invoke/[a-z0-9/\\?&=_A-Z]+")
	//matchString := compile.MatchString("/invoke/fff/ffff")
	//log.Println(matchString)
}
