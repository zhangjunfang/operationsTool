package shellbox

import (
	"github.com/astaxie/beego"
)

type QuestionsController struct {
	beego.Controller
}

func (this *QuestionsController) Get() {
	this.TplName = "shellbox/questions.html"
}
