package device

import (
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/zhangjunfang/operationsTool/soft"
)

// 内存信息
/**
{
	mem_total:#总内存
	mem_free:#剩余内存
}
*/
func RamInfo(user, password, ip string) map[string]string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RamInfo(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("cat /proc/meminfo | grep -E 'MemT|MemF';exit")
	temp := out.String()
	temparr1 := strings.Split(temp, "\n")
	var temparr2 = make([]string, 2)
	i := 0
	for i < 2 {
		temparr2[i] = strings.TrimSpace(strings.Split(temparr1[i], ":")[1])
		// fmt.Println(temparr2[i])
		i++
	}
	/*占用比*/
	temp1, err := strconv.Atoi(strings.Split(temparr2[1], " ")[0])
	temp2, err := strconv.Atoi(strings.Split(temparr2[0], " ")[0])
	ram_pre := 100 - float64(temp1)/float64(temp2)*100
	// 64位浮点数转字符串后两位
	ram_pre_str := strconv.FormatFloat(ram_pre, 'f', 2, 32)
	return map[string]string{"ram_total": temparr2[0], "ram_free": temparr2[1], "ram_pre": ram_pre_str + "%"}
}

func RamProcess(user, password, ip string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RamInfo(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("ps -ef")
	temp := out.String()
	return temp
}

func RamKill(user, password, ip, pid string) {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			RamKill(user, password, ip, pid)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	session.Run("kill " + pid)
}
