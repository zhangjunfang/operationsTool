package soft

import (
	"bytes"
	"fmt"
	"log"
)

func UserList(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			UserList(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("cat /etc/passwd | cut -f 1 -d :")
	fmt.Println(out.String())
	return out.String()
}

func UserDel(user, password, ip, auth string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			UserDel(user, password, ip, auth)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	var out2 bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out2
	session.Run("userdel " + auth)
	fmt.Println(out.String())
	fmt.Println(out2.String())
	if out2.String() != "" {
		return out2.String()
	} else {
		return out.String()
	}
}

func UserAdd(user, password, ip, auth, passwd string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			UserAdd(user, password, ip, auth, passwd)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	var out2 bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out2
	session.Run(fmt.Sprintf("useradd %s | echo %s | passwd --stdin %s", auth, passwd, auth))
	fmt.Println(out.String())
	fmt.Println(out2.String())
	if out2.String() != "" {
		return out2.String()
	} else {
		return out.String()
	}
}

func UserMod(user, password, ip, auth, group string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			UserMod(user, password, ip, auth, group)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	var out2 bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out2
	session.Run(fmt.Sprintf("usermod -g %s %s", group, auth))
	fmt.Println(out.String())
	fmt.Println(out2.String())
	if out2.String() != "" {
		return out2.String()
	} else {
		return out.String()
	}
}

func GroupList(user, password, ip string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			GroupList(user, password, ip)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	session.Stdout = &out
	session.Run("cat /etc/group | cut -f 1 -d :")
	fmt.Println(out.String())
	return out.String()
}

func GroupDel(user, password, ip, group string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			GroupDel(user, password, ip, group)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	var out2 bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out2
	session.Run("groupdel " + group)
	fmt.Println(out.String())
	fmt.Println(out2.String())
	if out2.String() != "" {
		return out2.String()
	} else {
		return out.String()
	}
}

func GroupAdd(user, password, ip, group string) string {
	session, err := Ssh(user, password, ip, 22)
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println("连接超时，意外中断")
			GroupAdd(user, password, ip, group)
		}
	}()
	if err != nil {
		log.Panic(err)
		return fmt.Sprintln(err)
	}
	var out bytes.Buffer
	var out2 bytes.Buffer
	session.Stdout = &out
	session.Stderr = &out2
	session.Run("groupadd " + group)
	fmt.Println(out.String())
	fmt.Println(out2.String())
	if out2.String() != "" {
		return out2.String()
	} else {
		return out.String()
	}
}
