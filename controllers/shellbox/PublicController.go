package shellbox

import (
	"github.com/astaxie/beego"
)

type PublicController struct {
	beego.Controller
}

func (this *PublicController) Get() {
	this.TplName = "shellbox/public.html"
}
