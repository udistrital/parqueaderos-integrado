package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:PropietarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseOne/controllers:VehiculoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
