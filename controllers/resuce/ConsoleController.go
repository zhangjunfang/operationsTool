package resuce

import (
	"github.com/astaxie/beego"
)

type ConsoleController struct {
	beego.Controller
}

func (this *ConsoleController) Get() {
	this.TplName = "resuce/webconsole.html"
}
