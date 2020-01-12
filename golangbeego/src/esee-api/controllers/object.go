package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200
// @Failure 403
// @router /eiam [get]
func (o *ObjectController) GetAll() {
	data := o.Ctx.GetCookie("eiamdata")
	o.Data["json"] = data
	o.Ctx.SetCookie("eiamdata", data, -1)
	o.ServeJSON()
}
