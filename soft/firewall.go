package soft

import (
	"bytes"
	"fmt"
	"log"
)

func IptablesCode(user, password, ip, code, auto string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			IptablesCode(user, password, ip, code, auto)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	var errs bytes.Buffer
	session.Stderr = &errs
	session.Run(code)
	fmt.Println(out.String())
	fmt.Println(errs.String())
	if auto == "true" {
		return IptablesSaveRestart(user, password, ip)
	} else {
		if errs.String() != "" {
			return errs.String()
		} else {
			return out.String()
		}
	}
}

func IptablesSaveRestart(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			IptablesSaveRestart(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
		fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("/etc/rc.d/init.d/iptables save;service iptables restart")
	fmt.Println(out.String())
	return out.String()
}

func IptablesList(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			IptablesList(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("iptables -L --line-numbers ")
	fmt.Println(out.String())
	return out.String()
}

func IptablesDel(user, password, ip, chain, num string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			IptablesDel(user, password, ip, chain, num)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run(fmt.Sprintf("iptables -D %s %s", chain, num))
	fmt.Println(out.String())
	return out.String()
}
