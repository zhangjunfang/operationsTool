package device

import (
	"github.com/zhangjunfang/operationsTool/device"
	"github.com/zhangjunfang/operationsTool/soft"

	"github.com/astaxie/beego"
)

type RamController struct {
	beego.Controller
}

func (this *RamController) Get() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["process"] = device.RamProcess(user, password, ip)
	soft.Notty(user, password, ip)
	this.TplName = "ram/ram-index.html"
}

func (this *RamController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = device.RamInfo(user, password, ip)
	this.ServeJSON()
}

func (this *RamController) Post() {
	pid := this.GetString("pid")
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	device.RamKill(user, password, ip, pid)
	this.Data["json"] = "success"
	this.ServeJSON()
}
