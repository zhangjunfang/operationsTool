package safe

import (
	"github.com/zhangjunfang/operationsTool/soft"

	"github.com/astaxie/beego"
)

type GroupController struct {
	beego.Controller
}

func (this *GroupController) Get() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.GroupList(user, password, ip)
	this.ServeJSON()
}

func (this *GroupController) Post() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	group := this.GetString("group")
	this.Data["json"] = soft.GroupAdd(user, password, ip, group)
	this.ServeJSON()
}

/*func (this *GroupController) Options() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
}

func (this *GroupController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
}*/

func (this *GroupController) Delete() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	group := this.GetString("group")
	this.Data["json"] = soft.GroupDel(user, password, ip, group)
	this.ServeJSON()
}
