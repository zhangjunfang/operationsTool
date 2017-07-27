package device

import (
	"github.com/zhangjunfang/operationsTool/device"

	"github.com/astaxie/beego"
)

type RomController struct {
	beego.Controller
}

func (this *RomController) Get() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["Romstatus"] = device.RomStatus(user, password, ip)
	this.TplName = "rom/rom-index.html"
}

func (this *RomController) Post() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	info := this.GetString("info")
	dev := this.GetString("dev")
	local := this.GetString("local")
	if local != "" {
		this.Data["json"] = map[string]string{"info": device.RomUmount(user, password, ip, local)}
	} else {
		this.Data["json"] = map[string]string{"info": device.RomMount(user, password, ip, dev, info)}
	}
	this.ServeJSON()
}

func (this *RomController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = device.RomInfo(user, password, ip)
	this.ServeJSON()
}

/*func (this *RomController) Delete() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	local := this.GetString("local")
	fmt.Println(local)
	this.Data["json"] = map[string]string{"info": device.RomUmount(user, password, ip, local)}
	this.ServeJSON()
}
*/
