package soft

import (
	"bytes"
	"fmt"
	"log"
)

func StopPing(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			StopPing(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all;cat /proc/sys/net/ipv4/icmp_echo_ignore_all")
	fmt.Println(out.String())
	return out.String()
}

// echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all;cat /proc/sys/net/ipv4/icmp_echo_ignore_all

func StartPing(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			StopPing(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("echo 0 > /proc/sys/net/ipv4/icmp_echo_ignore_all;cat /proc/sys/net/ipv4/icmp_echo_ignore_all")
	fmt.Println(out.String())
	return out.String()
}
