package controllers

import (
	"Parqueaderos/models"
	"encoding/json"
	//	"errors"
	"strconv"

	"github.com/astaxie/beego"
)

// operations for Ingreso
type IngresoController struct {
	beego.Controller
}

func (c *IngresoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Ingreso
// @Param	body		body 	models.Ingreso	true		"body for Ingreso content"
// @Success 201 {object} models.Ingreso
// @Failure 403 body is empty
// @router / [post]
func (c *IngresoController) Post() {

}

// @Title Get
// @Description get Ingreso by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ingreso
// @Failure 403 :id is empty
// @router /:id [get]
func (c *IngresoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id64, _ := strconv.Atoi(idStr)
	id := int16(id64)
	v, err := models.GetVehiculoByIdNfc(id)
	u, err1 := models.GetIslaByOcu(false)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	} else {
		if err1 != nil {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err1.Error()
		} else {
			z := models.Registro{IdVehiculo: v, IdIsla: u}
			if _, err := models.AddRegistro(&z); err == nil {
				c.Ctx.Output.SetStatus(201)
				key := []byte("LKHlhb899Y09olUi")
				aux, _ := json.Marshal(z)
				x, _ := models.Encrypt(key, aux)
				c.Data["json"] = x
			} else {
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = err.Error()
			}
		}
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Ingreso
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Ingreso
// @Failure 403
// @router / [get]
func (c *IngresoController) GetAll() {

}

// @Title Update
// @Description update the Ingreso
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Ingreso	true		"body for Ingreso content"
// @Success 200 {object} models.Ingreso
// @Failure 403 :id is not int
// @router /:id [put]
func (c *IngresoController) Put() {

}

// @Title Delete
// @Description delete the Ingreso
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *IngresoController) Delete() {

}
