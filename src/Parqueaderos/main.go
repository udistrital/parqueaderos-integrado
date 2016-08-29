package main

import (
	_ "Parqueaderos/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Conexion struct {
	database_host string
	database_port	string
	database_name	string
	database_user	string
	database_password	string
}

func init() {
	data, err := ioutil.ReadFile("conexion.yml")
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[interface{}]string)

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var database_host string = m["database_host"]
	var database_port string = m["database_port"]
	var database_name string = m["database_name"]
	var database_user string = m["database_user"]
	var database_password string = m["database_password"]

	orm.RegisterDataBase("default", "postgres",
		"postgres://"+database_user+":"+database_password+"@"+database_host+":"+database_port+"/"+database_name+"?sslmode=disable")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
