package safe

import (
	"github.com/zhangjunfang/operationsTool/soft"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.UserList(user, password, ip)
	this.ServeJSON()
}

func (this *UserController) Post() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	auth := this.GetString("auth")
	passwd := this.GetString("passwd")
	this.Data["json"] = soft.UserAdd(user, password, ip, auth, passwd)
	this.ServeJSON()
}

func (this *UserController) Options() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	auth := this.GetString("auth")
	group := this.GetString("group")
	this.Data["json"] = soft.UserMod(user, password, ip, auth, group)
	this.ServeJSON()
}

/*func (this *UserController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
}*/

func (this *UserController) Delete() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	auth := this.GetString("auth")
	this.Data["json"] = soft.UserDel(user, password, ip, auth)
	this.ServeJSON()
}
