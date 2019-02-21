package routers

import (
	"github.com/astaxie/beego"
	"search/container/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/resources", &controllers.ResourceController{})
	beego.Router("/resources/:id", &controllers.ResourceController{}, "put:Put")
}
