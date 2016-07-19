package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:IslaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:PropietarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["ReleaseTwo/controllers:VehiculoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
