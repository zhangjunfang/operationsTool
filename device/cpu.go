package device

import (
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/zhangjunfang/operationsTool/soft"
)

// CPU信息
/**
{
cpu_MHz:#Cpu频率
cpu_cores:#Cpu物理核心数
cpu_name:#Cpu型号
cpu_processor:#Cpu线程
}
*/
func CpuInfo(user, password, ip string) map[string]string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断,重新连接中!!!")
			CpuInfo(user, password, ip) //这里需要优化
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("cat /proc/cpuinfo | grep -i cpu;top -bn 1 | grep Cpu")
	temp := out.String()
	temp = strings.Replace(temp, " ", "", -1)
	temp = strings.Replace(temp, "\t", "", -1)
	var temparr1 []string
	temparr1 = strings.Split(temp, "\n")
	var temparr2 = make([]string, len(temparr1)) //Index out of range 处理方法
	for i := 0; i < (len(temparr1) - 1); i++ {
		temparr2[i] = strings.Split(temparr1[i], ":")[1]
		// fmt.Println(strings.Split(temparr1[i], ":")[1])
	}
	// cpu_pre := CpuPre(user, password, ip)
	return map[string]string{"cpu_name": temparr2[1], "cpu_MHz": temparr2[2], "cpu_cores": temparr2[3], "cpu_processor": strconv.Itoa((len(temparr1) - 2) / 5), "cpu_pre": temparr2[len(temparr1)-2]}
	// fmt.Println(temparr2)
	// return map[string]string{}
}

func CpuPre(user, password, ip string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			CpuPre(user, password, ip)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("top -bn 1 | grep Cpu")
	return out.String()
}

func CpuCore(user, password, ip string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			CpuCore(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("ls /sys/devices/system/cpu | grep cpu")
	temp := out.String()
	temparr := strings.Split(temp, "\n")
	return strconv.Itoa(len(temparr) - 3)
}

func CpuCoreStatus(user, password, ip string) string {
	session, err := soft.Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			CpuCoreStatus(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("cat /sys/devices/system/cpu/cpu*/online")
	temp := out.String()
	temp2 := strings.Replace(temp, "\n", "", -1)
	temp3 := strings.Replace(temp2, " ", "", -1)
	return temp3
}
