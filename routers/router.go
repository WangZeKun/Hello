package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WelcomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/change", &controllers.ChangeController{})
	beego.Router("/activities",&controllers.ActivityController{})
}
