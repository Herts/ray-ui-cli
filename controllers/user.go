package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
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