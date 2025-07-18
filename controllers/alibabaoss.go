package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"

	"alibabaoss/liboss"
)

type AlibabaOssController struct{ beego.Controller }

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ContentBody struct {
	Data string `json:data`
}

func (c *AlibabaOssController) PutString() {

	var content ContentBody
	json.Unmarshal(c.Ctx.Input.RequestBody, &content) //{data: "I am doggy"}

	if content.Data == "" {
		c.Data["json"] = Response{Success: false, Message: "Invalid Input", Error: "Body cannot be empty"}
	} else {
		if err := liboss.PutString("media/content/test.json", content.Data); err != nil {
			c.Data["json"] = Response{Success: false, Message: "Parameter Error", Error: err.Error()}
		} else {
			c.Data["json"] = Response{Success: true, Message: "Success"}
		}
	}

	c.ServeJSON()
}

func (c *AlibabaOssController) GetString() {

	if content, err := liboss.GetString("media/content/test.json"); err != nil {
		c.Data["json"] = Response{Success: false, Message: "Parameter Error", Error: err.Error()}
	} else {
		c.Data["json"] = Response{Success: true, Message: "Success", Data: content}
	}

	c.ServeJSON()
}

func (c *AlibabaOssController) DeleteStringFile() {

	err := liboss.Delete("media/content/test.json")
	if err != nil {
		c.Data["json"] = Response{Success: false, Message: "Failed to delete", Error: err.Error()}
	} else {
		c.Data["json"] = Response{Success: true, Message: "Success"}
	}

	c.ServeJSON()
}
