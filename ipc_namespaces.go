package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// sh-4.2# ipcs -q

// ------ Message Queues --------
// key        msqid      owner      perms      used-bytes   messages
// sh-4.2# ipcmk -Q
// Message queue id: 0
// sh-4.2# ipcs -q

// ------ Message Queues --------
// key        msqid      owner      perms      used-bytes   messages
// 0x7c0b410a 0          root       644        0            0

// 重新再使用go run ipc_namespaces.go,再创建一个新的sh环境新的namespaces,重新运行ipcs -q发现没有，则确定可以进行IPC隔离
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
