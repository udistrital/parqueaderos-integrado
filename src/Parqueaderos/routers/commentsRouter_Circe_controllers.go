package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
