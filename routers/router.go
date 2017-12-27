package routers

import (
	"github.com/ajaygh/genseq/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
