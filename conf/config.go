package conf

import (
	beego "github.com/beego/beego/v2/server/web"
)

var (
	OPENAI_API_KEY, _ = beego.AppConfig.String("openai::apiKey")
)
