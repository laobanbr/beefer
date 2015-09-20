package main

import (
	"github.com/astaxie/beego"
)

type BeeferController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

type User struct {
	Username string
}

// The Get method to handle the GET request
func (c *BeeferController) Get() {
	name := c.GetString("user")
	user := User{Username: name}
	c.Data["User"] = user

	c.TplNames = "beefer.tpl"
}

func (c *UserController) Login() {
	c.TplNames = "user/login.tpl"
}

func (c *UserController) Signup() {
	c.TplNames = "user/signup.tpl"
}

func (c *UserController) Logout() {
	c.TplNames = "beefer.tpl"
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Router("/user/login", &UserController{}, "get:Login")
	beego.Router("/user/signup", &UserController{}, "get:Signup")
	beego.Router("/user/logout", &UserController{}, "post:Logout")
	beego.Run(":8085")
}
