package soft

import (
	"github.com/zhangjunfang/operationsTool/soft"

	"github.com/astaxie/beego"
)

type SysController struct {
	beego.Controller
}

func (this *SysController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.SystemInfo(user, password, ip)
	this.ServeJSON()
}
