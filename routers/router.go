package routers

import (
	"user/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{},"*:Login")
	beego.Router("/index", &controllers.UserController{},"*:Index")
	beego.Router("/getAll", &controllers.UserController{},"get:GetAll")
	beego.Router("/delete/:id", &controllers.UserController{},"*:Delete")
	beego.Router("/form_add", &controllers.UserController{},"get:FormAdd")
	beego.Router("/post", &controllers.UserController{},"post:Post")
    beego.Router("/update/:id",&controllers.UserController{},"*:Update")
	beego.Router("/updateUser",&controllers.UserController{},"*:UpdateUser")
	beego.Router("/login",&controllers.UserController{},"*:Login")

	beego.Router("/loginTest",&controllers.UserController{},"*:LoginTest")
    beego.Router("/logout",&controllers.UserController{},"*:Logout")

}
