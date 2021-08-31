package service

import (
	"BullMock/models"
	"github.com/traefik/yaegi/interp"
)

func ExecScript(scriptName string, param ...interface{}) interface{} {
	script, err := models.GetByName(scriptName)
	if err == nil {
		inter := interp.New(interp.Options{})
		_, err := inter.Eval(script.ScriptContent)
		if err != nil {
			panic(err)
		}
		//v.
	}
	return err
}
