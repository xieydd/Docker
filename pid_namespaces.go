package main

//我们看见在环境中执行
//$pstree -pl 中pidnamespaces的pid为25275，但是namespaceah环境下执行echo $$为1
//─sshd(22136)───sshd(22140)───zsh(22141)───su(23606)───bash(23610)───go(25246)─┬─pid_namespaces(25275)─┬─sh(25279)──
import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
