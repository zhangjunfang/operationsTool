package safe

import (
	"github.com/astaxie/beego"
)

type UsergroupController struct {
	beego.Controller
}

func (this *UsergroupController) Get() {
	this.TplName = "safe/usergroup.html"
}
