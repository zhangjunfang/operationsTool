package file

import (
	"github.com/astaxie/beego"
)

type FilesystemController struct {
	beego.Controller
}

func (this *FilesystemController) Get() {
	this.TplName = "file/filesystem.html"
}
