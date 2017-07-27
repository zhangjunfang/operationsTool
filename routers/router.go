package routers

import (
	"github.com/zhangjunfang/operationsTool/controllers"
	"github.com/zhangjunfang/operationsTool/controllers/device"
	"github.com/zhangjunfang/operationsTool/controllers/file"
	"github.com/zhangjunfang/operationsTool/controllers/resuce"
	"github.com/zhangjunfang/operationsTool/controllers/safe"
	"github.com/zhangjunfang/operationsTool/controllers/shellbox"
	"github.com/zhangjunfang/operationsTool/controllers/soft"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.SshController{})
	beego.Router("/device/cpu", &device.CpuController{})
	beego.Router("/device/ram", &device.RamController{})
	beego.Router("/device/rom", &device.RomController{})
	beego.Router("/soft/system", &soft.SysController{})
	beego.Router("/resuce/webconsole", &resuce.ConsoleController{})
	beego.Router("/resuce/reboot", &resuce.RebootController{})
	beego.Router("/shellbox/public", &shellbox.PublicController{})
	beego.Router("/shellbox/private", &shellbox.PrivateController{})
	beego.Router("shellbox/questions", &shellbox.QuestionsController{})
	beego.Router("/safe/firewall", &safe.FirewallController{})
	beego.Router("/safe/stopping", &safe.StoppingController{})
	beego.Router("/safe/usergroup", &safe.UsergroupController{})
	beego.Router("/safe/user", &safe.UserController{})
	beego.Router("/safe/group", &safe.GroupController{})
	beego.Router("/file/filesystem", &file.FilesystemController{})
}
