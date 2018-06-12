package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// 我们发现在宿主机环境下为root 1 ，在namespace环境下为1234
// [root@00-0c-29-4a-bf-33 Docker]# id
// uid=0(root) gid=0(root) groups=0(root)
// [root@00-0c-29-4a-bf-33 Docker]# go run user_namespaces.go
// sh-4.2$ id
// uid=1234 gid=1234 groups=1234
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      0,
				Size:        1,
			},
		},
	}

	//cmd.SysProcAttr.Credential = &syscall.Credential{
	//	Uid: uint32(1),
	//	Gid: uint32(1),
	//}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(-1)
}
