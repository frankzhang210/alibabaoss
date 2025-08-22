package routers

import (
	"alibabaoss/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/oss",
			beego.NSRouter("/string", &controllers.AlibabaOssController{}, "post:PutString;get:GetString;delete:DeleteStringFile"),
		),
		beego.NSNamespace("/product",
			beego.NSRouter("", &controllers.ProductController{}, "post:InsertOrUpdate"),
			beego.NSRouter("/:product_id", &controllers.ProductController{}, "get:Get"),
		),
	)
	beego.AddNamespace(ns)
}
