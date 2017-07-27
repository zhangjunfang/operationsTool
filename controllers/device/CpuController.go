package device

import (
	"github.com/zhangjunfang/operationsTool/device"

	"github.com/astaxie/beego"
)

type CpuController struct {
	beego.Controller
}

func (this *CpuController) Get() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	CpuCore := device.CpuCore(user, password, ip)
	this.Data["CpuCore"] = CpuCore
	this.TplName = "cpu/cpu-index.html"
}

func (this *CpuController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = device.CpuInfo(user, password, ip)
	this.ServeJSON()
}

func (this *CpuController) Options() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = device.CpuCoreStatus(user, password, ip)
	this.ServeJSON()
}
