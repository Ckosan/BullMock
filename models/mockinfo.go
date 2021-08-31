package models

import (
	"BullMock/dao"
	"BullMock/parse"
	"encoding/json"
	"time"
)

type MockIns struct {
	ID              int    `gorm:"AUTO_INCREMENT"`
	Method          string `gorm:"column:method;size:8"`
	Url             string `gorm:"size:255"`
	Request         string `gorm:"size:1024"`
	Response        string `gorm:"size:1024"`
	ReqContentType  string `gorm:"size:255"`
	RespContentType string `gorm:"size:255"`
	ForwardUrl      string `gorm:"size:255"`
	MockRule        string `gorm:"size:255"`
	Status          bool
	CreatedAt       time.Time
	UpdateAt        time.Time
}

func mockInstance2Ins(mock *parse.MockInstance) *MockIns {
	mockIns := MockIns{
		Method:          mock.Method,
		Url:             mock.Url,
		ReqContentType:  mock.ReqContentType,
		RespContentType: mock.RespContentType,
		ForwardUrl:      mock.ForwardUrl,
		MockRule:        mock.MockRule,
		CreatedAt:       time.Now(),
		UpdateAt:        time.Now(),
	}
	req, _ := json.Marshal(mock.Request)
	resp, _ := json.Marshal(mock.Response)
	mockIns.Request = string(req)
	mockIns.Response = string(resp)
	return &mockIns
}
func CreateMockIns(mock *parse.MockInstance) error {
	mockIns := mockInstance2Ins(mock)
	err := dao.DB.Create(&mockIns).Error
	return err
}

func UpdateMockIns(mock *parse.MockInstance) (err error) {
	mockIns := mockInstance2Ins(mock)
	err = dao.DB.Save(&mockIns).Error
	return
}

func GetMockIns(method, url string) (mock *MockIns, err error) {
	mock = new(MockIns)
	if err = dao.DB.Where("method=? and url=?", method, url).First(mock).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateMockStatus(status bool, method, url string) {
	mockInfo := new(MockIns)
	mockInfo.Method = method
	mockInfo.Url = url
	dao.DB.Model(&mockInfo).Update("status", status)
}

func DeleteMockIns(method, url string) (err error) {
	err = dao.DB.Where("method=? and url=?", method, url).Delete(&MockIns{}).Error
	return
}
