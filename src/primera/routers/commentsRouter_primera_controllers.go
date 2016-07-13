package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["primera/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["primera/controllers:MigrationsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["primera/controllers:MigrationsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["primera/controllers:MigrationsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["primera/controllers:MigrationsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["primera/controllers:MigrationsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:ObjectController"] = append(beego.GlobalControllerRouter["primera/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:ObjectController"] = append(beego.GlobalControllerRouter["primera/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:ObjectController"] = append(beego.GlobalControllerRouter["primera/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:ObjectController"] = append(beego.GlobalControllerRouter["primera/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:ObjectController"] = append(beego.GlobalControllerRouter["primera/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["primera/controllers:PropietarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["primera/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["primera/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["primera/controllers:PropietarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["primera/controllers:PropietarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:UserController"] = append(beego.GlobalControllerRouter["primera/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["primera/controllers:VehiculoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["primera/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["primera/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["primera/controllers:VehiculoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["primera/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["primera/controllers:VehiculoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
