package main

import (
	"BullMock/dao"
	"BullMock/models"
	"BullMock/service"
	"BullMock/service/pluginserver"
	"BullMock/setting"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	addMock      = "/mock/add"
	updateMock   = "/mock/update"
	closeMock    = "/mock/close"
	deleteMock   = "/mock/delete"
	addScript    = "/add/Script"
	updateScript = "/update/Script"
	invoke       = "/all"
)

func startMock() {
	r := pluginserver.NewRouter()
	r.HandleFunc(addMock, service.AddMock).Methods("POST")
	r.HandleFunc(updateMock, service.UpdateMock).Methods("POST")
	r.HandleFunc(closeMock, service.CloseMock).Methods("GET")
	r.HandleFunc(deleteMock, service.DeleteMock).Methods("GET")
	r.HandleFunc(addScript, service.AddScript).Methods("POST")
	r.HandleFunc(updateScript, service.UpScript).Methods("POST")
	r.HandleFunc(invoke, service.Accept)
	log.Fatal(http.ListenAndServe(":9091", r))
}

func main() {
	logInit()
	//加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		log.Printf("load config from file failed, err:%v\n", err)
		return
	}
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		log.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	dao.DB.AutoMigrate(&models.MockIns{})
	dao.DB.AutoMigrate(&models.Script{})
	startMock()
}

var logger *log.Logger

func logInit() {
	file := "./log/" + "bull_" + time.Now().Format("20180102") + "_log.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	logger = log.New(logFile, "[bull]", log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出
	return

}
