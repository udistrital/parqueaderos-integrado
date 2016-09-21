package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/session"
)

// operations for Login
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	//c.EnableXSRF = false
}
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// @Title Post
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
}

// @Title Get
// @Description get Login by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Login
// @Failure 403 :id is empty
// @router / [get]
func (c *LoginController) Get() {
	username := c.Ctx.Input.Query("username")
	password := c.Ctx.Input.Query("password")

	tokenString := false
	if username == "admin" && password == "admin" {
		tokenString = true
		//	c.EnableXSRF = true
		c.Data["json"] = "ok"

	}
	fmt.Println(tokenString)
	c.ServeJSON(tokenString)

}
func (c *LoginController) ServeJSON(chkuser bool, encoding ...bool) {
	var (
		hasIndent   = true
		hasEncoding = false
	)
	//	if BConfig.RunMode == PROD {
	//		hasIndent = false
	//	}
	if len(encoding) > 0 && encoding[0] == true {
		hasEncoding = true
	}
	c.Ctx.Output.JSON(c.Data["json"], hasIndent, hasEncoding, chkuser)
}
