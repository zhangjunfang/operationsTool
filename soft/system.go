package soft

import (
	"bytes"
	"log"
)

func SystemInfo(user, password, ip string) map[string]string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			SystemInfo(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("cat /etc/issue | head -n 1")
	// fmt.Println(out.String())
	return map[string]string{"system_info": out.String()}
}
