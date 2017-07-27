package device

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/zhangjunfang/operationsTool/soft"
)

func RomInfo(user, password, ip string) map[int]string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RomInfo(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("df -h;exit")
	temp := out.String()
	temparr1 := strings.Split(temp, "\n")
	/*	map如此初始化才能消除assignment to entry in nil map错误*/
	tempmap := make(map[int]string)
	var templist string
	for i := 1; i < len(temparr1)-1; i++ {
		temparr2 := strings.Split(temparr1[i], " ")
		for _, val := range temparr2 {
			if val != "" {
				templist = templist + val + ","
			}
		}
		tempmap[i-1] = templist
		templist = ""
	}
	return tempmap
}

func RomStatus(user, password, ip string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RomStatus(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("fdisk -l | grep Disk | cut -f 1 -d :| grep dev;echo '#';mount -l | grep dev | awk '{print $1\" \"$3}'")
	// fmt.Println(out.String())
	return out.String()
}

func RomMount(user, password, ip, dev, info string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RomMount(user, password, ip, dev, info)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	err = session.Run(fmt.Sprintf("mount %s %s", dev, info))
	if err != nil {
		return fmt.Sprintln(err)
	} else {
		return "success"
	}
}

func RomUmount(user, password, ip, local string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RomUmount(user, password, ip, local)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	err = session.Run(fmt.Sprintf("umount %s", local))
	if err != nil {
		return fmt.Sprintln(err)
	} else {
		return "success"
	}
}
