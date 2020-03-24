package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type UserController struct {
	beego.Controller
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (c *UserController) CreateUser() {
	var newUser models.User

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newUser)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}
	newUser.Email = strings.ReplaceAll(newUser.Email, " ", "")
	u := models.GetUser(newUser.Email)
	if u.Email == newUser.Email {
		c.Data["json"] = response{Message: fmt.Sprintf("User email %s exists", newUser.Email)}
		c.ServeJSON()
		return
	}

	if len(newUser.UserId) != 36 {
		c.Data["json"] = response{
			Message: "User id is not correct",
		}
		c.ServeJSON()
		return
	}

	models.AddUser(&newUser)
	c.Data["json"] = response{Message: "User creation success"}
	c.ServeJSON()
}

func (c *UserController) UpdateUser() {
	var user models.User

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		c.Data["json"] = response{Message: err.Error()}
		c.ServeJSON()
		return
	}

	u := models.GetUser(user.Email)
	if u.Email != user.Email {
		c.Data["json"] = response{Message: fmt.Sprintf("User email %s does not exist", user.Email)}
		c.ServeJSON()
		return
	}

	if len(user.UserId) != 36 {
		c.Data["json"] = response{
			Message: "User id is not correct",
		}
		c.ServeJSON()
		return
	}

	user.Model = u.Model

	models.UpdateUser(&user)
	c.Data["json"] = response{Message: "User update success"}
	c.ServeJSON()
}

func (c *UserController) DeleteUser() {

}

func (c *UserController) ListAllUsers() {
	users := models.GetAllUser()
	c.Data["json"] = response{Data: users}
	c.ServeJSON()
}

func (c *UserController) ListAllDataConsumed() {
	c.Data["json"] = response{Message: "success", Data: models.GetAllDataConsumed()}
	c.ServeJSON()
}

func (c *UserController) UpdateDataConsumed() {
	emails := models.UpdateDataConsumed()
	c.Data["json"] = response{Message: "success", Data: emails}
	c.ServeJSON()
}

func (c *UserController) GetUserDataPage() {
	data := models.GetAllDataConsumed()
	// TODO: the form to display data should be changed
	for _, d := range data {
		d.UpDataConsumed /= 10e6
		d.DownDataConsumed /= 10e6
	}
	c.Data["data"] = data
	c.Layout = "layout.html"
	c.TplName = "userdata.html"
}

func (c *UserController) GetUserPage() {
	//models.UpdateDataConsumedInUser()
	users := models.GetAllUser()
	c.Data["data"] = users
	c.Layout = "layout.html"
	c.TplName = "user.html"
}
