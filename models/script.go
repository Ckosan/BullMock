package models

import (
	"BullMock/dao"
	"time"
)

type Script struct {
	ScriptContent string `gorm:"type:text"`
	ScripType     string `gorm:"size:64"`
	ScriptName    string `gorm:"size:128"`
	CreatedAt     time.Time
	UpdateAt      time.Time
}

func AddScript(script *Script) error {
	err := dao.DB.Create(&script).Error
	return err
}

func UpScript(script *Script) error {
	err := dao.DB.Save(&script).Error
	return err
}

func GetByName(name string) (*Script, error) {
	script := new(Script)
	if err := dao.DB.Where("script_name=?", name).First(script).Error; err != nil {
		return nil, err
	}
	return script, nil
}
