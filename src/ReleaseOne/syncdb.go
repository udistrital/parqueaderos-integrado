package main

import (
	_ "ReleaseOne/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:grupoglud=@tcp(127.0.0.1:3306)/ReleaseOne")
}

func main() {
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}
}

