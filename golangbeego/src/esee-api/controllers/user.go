package controllers

import (
	"context"
	"esee-api/models"
	"esee-api/util"
	"fmt"
	"gopkg.in/olivere/elastic.v5"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description 测试es
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {

	var res *elastic.SearchResult
	var err error
	query := elastic.NewQueryStringQuery("user:Jame3")
	res, err = util.Client.Search("esee").
		Type("ci").
		Sort("age", true).
		Size(20).
		Query(query).
		Pretty(true).
		SearchAfter("329").
		Do(context.Background())
	fmt.Println(err)
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
