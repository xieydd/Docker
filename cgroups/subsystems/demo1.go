package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"syscall"
	"path"
	"io/ioutil"
	"strconv"
)

//这里发现并没有限制，还没找出错误

//挂载memory subsystem的hierarchy的根目录文件
const cgroupMemoryHierarchMount = "/sys/fs/cgroup/memory"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("Current pid %d", syscall.Getpid())
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Err", err)
			log.Fatal(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println("Err", err)
		os.Exit(1)
	}

	//得到fork出来的进程映射到外面命名空间的pid
	fmt.Printf("%v",cmd.Process.Pid)
	//在系统默认创建了memory subsystem hierarchy上创建cgroup
	os.Mkdir(path.Join(cgroupMemoryHierarchMount,"testmemorylimit"),0755)
	//将容器进程加入到这个cgroup中
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchMount,"testmemorylimit","tasks"),[]byte(strconv.Itoa(cmd.Process.Pid)),0644)	//限制cgroup进程使用
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchMount,"testmemorylimit","memory.limit_in_bytes"),[]byte("100m"),0644)
	cmd.Process.Wait()

}
