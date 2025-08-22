package controllers

import (
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

var (
	runMode string
)

type BaseController struct{ beego.Controller }

func init() {
	runMode, _ = beego.AppConfig.String("runmode")
}

func (c *BaseController) GetInputParamNoValidation(name string) string {

	name = strings.TrimLeft(name, ":")
	param := strings.TrimSpace(c.Ctx.Input.Param(fmt.Sprintf(":%s", name)))

	return param
}

func (c *BaseController) ResponseMessage(status bool, message interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
