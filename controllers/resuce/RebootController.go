package resuce

import (
	"github.com/astaxie/beego"
	"github.com/zhangjunfang/operationsTool/soft"
)

type RebootController struct {
	beego.Controller
}

func (this *RebootController) Get() {
	this.TplName = "resuce/reboot.html"
}

func (this *RebootController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.Reboot(user, password, ip)
	this.DelSession("user")
	this.DelSession("password")
	this.DelSession("ip")
	this.ServeJSON()
}
