package main

// 这里启动 go run uts_namespaces.go后进入一个sh的运行环境
// yum install psmisc && pstree -pl
// ─sshd(22136)───sshd(22140)───zsh(22141)───su(23606)───bash(23610)───go(23769)─┬─uts_namespaces(23797)─┬─sh(23801)
// 发现父进程和子进程的UTS-namespace不同
// [root@15-pxe Docker]# readlink /proc/23797/ns/uts
// uts:[4026531838]
// [root@15-pxe Docker]# readlink /proc/23801/ns/uts
// uts:[4026532163]
//我们发现UTS Namespaces可以对namespaces进行隔离
// sh-4.2# hostname -b bird
// sh-4.2# hostname
// bird
// [root@15-pxe Docker]# hostname
// 15-pxe

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("sh")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
