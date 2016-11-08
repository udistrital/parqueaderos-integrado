package main

import (
	"Parqueaderos/conf"
	_ "Parqueaderos/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	p := conf.Parameters
	orm.RegisterDataBase("default", "postgres",
		"postgres://"+p.DATABASE_USER+":"+p.DATABASE_PASSWORD+"@"+p.DATABASE_HOST+":"+p.DATABASE_PORT+"/"+p.DATABASE_NAME+"?sslmode=disable&search_path="+p.DATABASE_SCHEMA)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
