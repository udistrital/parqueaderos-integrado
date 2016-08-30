package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:GrupoIslaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:IslaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:PropietarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:TipoPropietarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"] = append(beego.GlobalControllerRouter["Parqueaderos/controllers:VehiculoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

}
