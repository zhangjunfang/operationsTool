package safe

import (
	"github.com/zhangjunfang/operationsTool/soft"

	"github.com/astaxie/beego"
)

type StoppingController struct {
	beego.Controller
}

func (this *StoppingController) Get() {
	this.TplName = "safe/stopping.html"
}

func (this *StoppingController) Delete() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.StopPing(user, password, ip)
	this.ServeJSON()
}

func (this *StoppingController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.StartPing(user, password, ip)
	this.ServeJSON()
}
