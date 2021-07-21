package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.facebook.com/profile.php?id=100039939467069"
	c.Data["Email"] = "moha.denta@gmail.com"
	c.Data["EmailName"] = "Mohamed Abdulfattah"
	c.TplName = "index.tpl"
}
func (c *MainController) SignUp() {
	c.TplName = "signup.tpl"
}
