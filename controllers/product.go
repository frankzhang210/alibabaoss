package controllers

import (
	m "alibabaoss/models"
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
)

type ProductController struct{ BaseController }

func (c *ProductController) InsertOrUpdate() {

	product := &m.Product{}
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(product); err != nil {
		c.Data["json"] = c.ResponseMessage(false, err.Error())
		c.ServeJSON()
		return
	}

	record, err2 := m.InsertOrUpdateProduct(product)
	if err2 != nil {
		logs.Error(err2)
		c.Data["json"] = c.ResponseMessage(false, err2.Error())
	} else {
		c.Data["json"] = c.ResponseMessage(true, record)
	}
	c.ServeJSON()
}

/*********************************************
*	Get
**********************************************/
func (c *ProductController) Get() {

	product_id := c.GetInputParamNoValidation("product_id")

	product, err := m.GetProduct(product_id)
	if err != nil {
		c.Data["json"] = c.ResponseMessage(false, err.Error())

		c.ServeJSON()
		c.StopRun()
		return
	}

	c.Data["json"] = c.ResponseMessage(true, product)
	c.ServeJSON()
}
