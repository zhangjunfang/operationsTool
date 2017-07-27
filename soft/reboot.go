package soft

import (
	"bytes"
	"log"
)

/*服务器重启*/
func Reboot(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			Reboot(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("reboot")
	// fmt.Println(out.String())
	return out.String()
}
