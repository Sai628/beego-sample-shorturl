package routers

import (
	"github.com/Sai628/beego-sample-shorturl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})
}
