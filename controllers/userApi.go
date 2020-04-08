package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Herts/ray-ui-cli/models"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type UserApiController struct {
	beego.Controller
}

// @Title ListAllDataConsumed
// @Description List all users' data
// @Param startDate query *time.Time false
// @Param endDate query *time.Time false
// @Success 200 {response} response
// @router /listData [get]
func (c *UserApiController) ListAllDataConsumed(startDate, endDate *time.Time) {
	data := models.GetAllDataConsumedInRange(startDate, endDate)
	c.Data["json"] = response{Message: "success", Data: data}
	c.ServeJSON()
}

// @Param body body models.User true "user"
// @Success 200 {response} response
// @router /add [post]
func (c *UserApiController) CreateUser() {
	var newUser models.User

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newUser)
	if err != nil {
		c.Data["json"] = response{Message: "json.Unmarshal" + err.Error()}
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
	c.Data["json"] = response{Message: fmt.Sprintf("User %s creation success", newUser.Email)}
	c.ServeJSON()
}

// @Param body body models.User true "user"
// @Success 200 {response} response
// @router /update [put]
func (c *UserApiController) UpdateUser() {
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

func (c *UserApiController) DeleteUser() {

}

// @Success 200 {response} response
// @router /list [get]
func (c *UserApiController) ListAllUsers() {
	users := models.GetAllUser()
	c.Data["json"] = response{Data: users}
	c.ServeJSON()
}

func (c *UserApiController) UpdateDataConsumed() {
	emails := models.UpdateDataConsumed()
	c.Data["json"] = response{Message: "success", Data: emails}
	c.ServeJSON()
}
