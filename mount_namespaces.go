package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

//我们看见在环境中执行
//$pstree -pl 中pidnamespaces的pid为25275，但是namespaceah环境下执行echo $$为1
//─sshd(22136)───sshd(22140)───zsh(22141)───su(23606)───bash(23610)───go(25246)─┬─pid_namespaces(25275)─┬─sh(25279)──

//我们可以注意到这个proc可以通过内核和内核模块将信息发送给进程
//这里的/proc文件是host的
// sh-4.2# cd /proc/
// sh-4.2# ls
// 1      15317  2      22136  23720  23906  26     329  37    465  51   677  699   774        dma      kpagecount    self
// 10     15318  2045   22140  23724  23907  27     33   377   468  52   678  7     79     driver   kpageflags    slabinfo
// 100    15575  2057   22141  23725  23911  27644  330  378   469  53   679  700   8      execdomains  loadavg       softirqs
// 10206  15799  2061   22203  23748  23912  28     331  379   47   54   680  719   81     fb       locks         stat
// 1076   16     20754  22325  23802  24029  28928  332  38    470  55   681  73    867        filesystems  mdstat        swaps
// 1080   1615   20770  22539  23806  24820  29806  333  380   471  553  682  7373  9      fs       meminfo       sys
// 11     16205  20848  22540  23807  24824  3  334  390   472  576  683  7378  acpi       interrupts   misc          sysrq-trigger
// 1121   16216  2086   22859  23856  24841  30     335  41    473  577  684  7397  buddyinfo  iomem    modules       sysvipc
// 12     17     20990  23     23860  24842  30248  336  42    474  61   685  7398  bus        ioports  mounts        timer_list
// 13     17372  21     23138  23872  24896  30252  337  43    475  62   686  741   cgroups    ipmi     mtrr          timer_stats
// 136    17376  21027  23142  23890  25287  30773  338  441   476  628  693  743   cmdline    irq      net           tty
// 14572  17397  21043  23159  23891  25629  30774  340  442   477  63   694  744   consoles   kallsyms     pagetypeinfo  uptime
// 14573  18     2124   23160  23895  25630  31     341  45    48   635  695  746   cpuinfo    kcore    partitions    version
// 15     18751  2125   23577  23896  25658  318    35   451   49   64   696  747   crypto     keys     sched_debug   vmallocinfo
// 15212  19883  21960  23606  23904  25662  32     350  452   5    65   697  76    devices    key-users    schedstat     vmstat
// 15267  19916  22     23610  23905  25663  322    36   4566  50   672  698  772   diskstats  kmsg     scsi          zoneinfo
//挂载到该namespace下后系统进程只有本namespace的，docker mount利用这个特性
// sh-4.2# mount -t proc proc /proc
// sh-4.2# ls /proc
// 1      cmdline    dma      interrupts  kcore       loadavg  mounts    schedstat  swaps      tty
// 4      consoles   driver       iomem       keys    locks    mtrr      scsi       sys        uptime
// acpi       cpuinfo    execdomains  ioports     key-users   mdstat   net       self       sysrq-trigger  version
// buddyinfo  crypto     fb       ipmi        kmsg    meminfo  pagetypeinfo  slabinfo   sysvipc        vmallocinfo
// bus    devices    filesystems  irq         kpagecount  misc     partitions    softirqs   timer_list     vmstat
// cgroups    diskstats  fs       kallsyms    kpageflags  modules  sched_debug   stat       timer_stats    zoneinfo
// sh-4.2# ps -ef
// UID        PID  PPID  C STIME TTY          TIME CMD
// root         1     0  0 09:21 pts/4    00:00:00 sh
// root         5     1  0 09:24 pts/4    00:00:00 ps -ef
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
