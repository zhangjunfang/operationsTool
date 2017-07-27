package controllers

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/astaxie/beego"
	"github.com/zhangjunfang/operationsTool/soft"
)

type SshController struct {
	beego.Controller
}

func (this *SshController) Get() {
	this.Data["LoginInfo"] = this.GetSession("LoginInfo")
	this.DelSession("LoginInfo")
	fmt.Println("Connecting:", this.GetSession("ip"))
	if this.GetSession("ip") != nil {
		user := this.GetSession("user").(string)
		password := this.GetSession("password").(string)
		ip := this.GetSession("ip").(string)
		soft.Notty(user, password, ip)
		this.TplName = "ssh/SshDetail.html"
	} else {
		this.TplName = "ssh/SshLogin.html"
	}
}

func (this *SshController) Post() {

	userip := this.GetString("userip")
	password := this.GetString("password")
	arr := strings.Split(userip, "@")
	user := arr[0]
	ip := arr[1]
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	/**
	ssh连接判断
	*/
	defer func() {
		if err := recover(); err != nil {
			log.Println("Connect Error")
			this.SetSession("LoginInfo", "error")
			this.Redirect("/", 302)
		}
	}()
	if err != nil {
		panic(err)
		log.Println(err)
	}
	/**
	错误恢复
	*/
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("echo Connect Success")
	if strings.Contains(out.String(), "Success") {
		this.SetSession("user", user)
		this.SetSession("ip", ip)
		this.SetSession("password", password)
		this.Data["LoginInfo"] = "success"
		this.Redirect("/", 302)
	} else {
		panic(err)
		log.Println("Connect Fail")
	}

}

// 断开连接
func (this *SshController) Delete() {
	this.DelSession("user")
	this.DelSession("ip")
	this.DelSession("password")
	this.Data["json"] = "success"
	this.ServeJSON()
}
