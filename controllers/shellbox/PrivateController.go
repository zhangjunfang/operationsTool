package shellbox

import (
	"github.com/astaxie/beego"
)

type PrivateController struct {
	beego.Controller
}

func (this *PrivateController) Get() {
	this.TplName = "shellbox/private.html"
}
