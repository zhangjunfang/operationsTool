package safe

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/zhangjunfang/operationsTool/soft"
)

type FirewallController struct {
	beego.Controller
}

func (this *FirewallController) Get() {
	this.TplName = "safe/firewall.html"
}

func (this *FirewallController) Post() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	code := this.GetString("code")
	auto := this.GetString("auto")
	if strings.Contains(code, "iptables") {
		res := soft.IptablesCode(user, password, ip, code, auto)
		if res == "" || strings.Contains(res, "OK") {
			this.Data["json"] = "执行成功"
		} else {
			this.Data["json"] = "执行失败，请重新检查规则。" + res
		}
	} else {
		this.Data["json"] = "非有效的iptables规则"
	}
	this.ServeJSON()
}

func (this *FirewallController) Put() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = strings.Replace(soft.IptablesList(user, password, ip), "\n", ";", -1)
	this.ServeJSON()
}

func (this *FirewallController) Options() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	this.Data["json"] = soft.IptablesSaveRestart(user, password, ip)
	this.ServeJSON()
}

func (this *FirewallController) Delete() {
	user := this.GetSession("user").(string)
	password := this.GetSession("password").(string)
	ip := this.GetSession("ip").(string)
	chain := this.GetString("chain")
	num := this.GetString("num")
	this.Data["json"] = soft.IptablesDel(user, password, ip, chain, num)
	this.ServeJSON()
}
