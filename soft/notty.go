package soft

import (
	"log"
)

/*清除积累的sshd ： root@notty进程*/
func Notty(user, password, ip string) {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			Notty(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
	}
	session.Run("ps -efww|grep root@notty|grep -v grep|cut -c 9-15|xargs kill -9")
}
