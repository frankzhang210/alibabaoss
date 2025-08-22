package models

import (
	"log"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	url, err := beego.AppConfig.String("mysql::url")
	if err != nil || url == "" {
		log.Fatalf("FATAL: Could not find 'mysql::url' or it is empty in app.conf. %v", err)
	}

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", url)
	orm.RunSyncdb("default", false, true)
}
