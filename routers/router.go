package routers

import (
	"github.com/Sai628/beego-sample-shorturl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
