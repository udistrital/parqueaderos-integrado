package main

import (
	_ "ReleaseTwo/docs"
	_ "ReleaseTwo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://olzambranop:123Abcd@127.0.0.1:5432/releasetwo?sslmode=disable")
	orm.RunCommand()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Content-Type", "Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))
	}
	beego.Run()
}
