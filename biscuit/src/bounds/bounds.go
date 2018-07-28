package bounds

import "runtime"

import "res"

type Boundkey_t int

const (
	B_SYS_PAUSE Boundkey_t = iota
	B_SYS_PREAD
	B_USERBUF_T__TX
	B_TCPFOPS_T_READ
	B_SYS_POLL
	B_SYS_WRITEV
	B_SYS_PROF
	B_USERIOVEC_T__TX
	B_IMEMNODE_T__DESCAN
	B_SYS_TRUNCATE
	B_SYS_SOCKET
	B_SYS_BIND
	B_SYS_CHDIR
	B_PROC_T_K2USER_INNER
	B_RAWDFOPS_T_WRITE
	B_RAWDFOPS_T_READ
	B_IMEMNODE_T_DO_WRITE
	B_IMEMNODE_T_IWRITE
	B_SYS_READV
	B_SYS_SENDTO
	B_SYS_RECVFROM
	B_SYS_SETRLIMIT
	B_INSERTARGS
	B_SYS_WRITE
	B_SYS_GETCWD
	B_TCPTIMERS_T__TCPTIMERS_DAEMON
	B_SYS_SOCKETPAIR
	B_SYS_LINK
	B_SYS_MKNOD
	B_SYS_GETTID
	B_PROC_T_USERARGS
	B_SYS_UNLINK
	B_SYS_REBOOT
	B_KBD_DAEMON
	B_SYS_FSTAT
	B_SYS_MUNMAP
	B_SYS_RECVMSG
	B_SYS_FCNTL
	B_SYS_RENAME
	B_ELF_T_ELF_LOAD
	B_SYS_FTRUNCATE
	B_SYS_NANOSLEEP
	B_SYS_PIPE2
	B_SYS_STAT
	B_SYS_LSEEK
	B_SYS_CONNECT
	B_SYS_SENDMSG
	B_SYS_GETSOCKOPT
	B_SYS_INFO
	B_PROC_T_USER2K_INNER
	B_TCPFOPS_T_WRITE
	B_SYS_GETTIMEOFDAY
	B_SYS_GETRUSAGE
	B_SYS_PWRITE
	B_FS_T__FULLPATH
	B_IMEMNODE_T_IREAD
	B_ARP_RESOLVE
	B_IMEMNODE_T_IMMAPINFO
	B_PIPE_T_OP_FDADD
	B_SYS_ACCESS
	B_SYS_LISTEN
	B_SYS_EXECV
	B_SYS_GETRLIMIT
	B_SYS_SYNC
	B_FUTEX_T_FUTEX_START
	B_SYS_MKDIR
	B_SYS_KILL
	B_SYS_FUTEX
	B_USERIOVEC_T_IOV_INIT
	B_SYS_OPEN
	B_SYS_MMAP
	B_SYS_DUP2
	B_SYS_ACCEPT
	B_SYS_SHUTDOWN
	B_FS_T__ISANCESTOR
	B_PIPEFOPS_T_WRITE
	B_SYS_GETPID
	B_SYS_READ
	B_SYS_THREXIT
	B_FS_T_FS_RENAME
	B_FS_T_FS_NAMEI
	B_IMEMNODE_T_BMAPFILL
	B_SYS_SIGACTION
	B_SYS_GETPPID
	B_SYS_SETSOCKOPT
	B_SYS_FORK
	B_SYS_WAIT4
	B_SYSCALL_T_SYS_CLOSE
	B_SYSCALL_T_SYS_EXIT
	B_IMEMNODE_T_IFREE
	B_BITMAP_T_APPLY
	B_ICACHE_T_FREEDEAD
	B_IXGBE_T_INT_HANDLER
	B_AHCI_PORT_T_QUEUEMGR
	B_AHCI_DISK_T_INT_HANDLER
)

func Bounds(k Boundkey_t) *res.Res_t {
	return boundres[k]
}

var boundres = []*res.Res_t {
	B_SYS_PAUSE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_PAUSE]))}},
	B_SYS_PREAD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_PREAD]))}},
	B_USERBUF_T__TX: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_USERBUF_T__TX]))}},
	B_TCPFOPS_T_READ: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_TCPFOPS_T_READ]))}},
	B_SYS_POLL: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_POLL]))}},
	B_SYS_WRITEV: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_WRITEV]))}},
	B_SYS_PROF: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_PROF]))}},
	B_USERIOVEC_T__TX: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_USERIOVEC_T__TX]))}},
	B_IMEMNODE_T__DESCAN: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T__DESCAN]))}},
	B_SYS_TRUNCATE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_TRUNCATE]))}},
	B_SYS_SOCKET: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SOCKET]))}},
	B_SYS_BIND: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_BIND]))}},
	B_SYS_CHDIR: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_CHDIR]))}},
	B_PROC_T_K2USER_INNER: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_PROC_T_K2USER_INNER]))}},
	B_RAWDFOPS_T_WRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_RAWDFOPS_T_WRITE]))}},
	B_RAWDFOPS_T_READ: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_RAWDFOPS_T_READ]))}},
	B_IMEMNODE_T_DO_WRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T_DO_WRITE]))}},
	B_IMEMNODE_T_IWRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T_IWRITE]))}},
	B_SYS_READV: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_READV]))}},
	B_SYS_SENDTO: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SENDTO]))}},
	B_SYS_RECVFROM: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_RECVFROM]))}},
	B_SYS_SETRLIMIT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SETRLIMIT]))}},
	B_INSERTARGS: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_INSERTARGS]))}},
	B_SYS_WRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_WRITE]))}},
	B_SYS_GETCWD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETCWD]))}},
	B_TCPTIMERS_T__TCPTIMERS_DAEMON: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_TCPTIMERS_T__TCPTIMERS_DAEMON]))}},
	B_SYS_SOCKETPAIR: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SOCKETPAIR]))}},
	B_SYS_LINK: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_LINK]))}},
	B_SYS_MKNOD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_MKNOD]))}},
	B_SYS_GETTID: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETTID]))}},
	B_PROC_T_USERARGS: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_PROC_T_USERARGS]))}},
	B_SYS_UNLINK: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_UNLINK]))}},
	B_SYS_REBOOT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_REBOOT]))}},
	B_KBD_DAEMON: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_KBD_DAEMON]))}},
	B_SYS_FSTAT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_FSTAT]))}},
	B_SYS_MUNMAP: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_MUNMAP]))}},
	B_SYS_RECVMSG: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_RECVMSG]))}},
	B_SYS_FCNTL: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_FCNTL]))}},
	B_SYS_RENAME: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_RENAME]))}},
	B_ELF_T_ELF_LOAD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_ELF_T_ELF_LOAD]))}},
	B_SYS_FTRUNCATE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_FTRUNCATE]))}},
	B_SYS_NANOSLEEP: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_NANOSLEEP]))}},
	B_SYS_PIPE2: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_PIPE2]))}},
	B_SYS_STAT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_STAT]))}},
	B_SYS_LSEEK: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_LSEEK]))}},
	B_SYS_CONNECT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_CONNECT]))}},
	B_SYS_SENDMSG: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SENDMSG]))}},
	B_SYS_GETSOCKOPT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETSOCKOPT]))}},
	B_SYS_INFO: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_INFO]))}},
	B_PROC_T_USER2K_INNER: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_PROC_T_USER2K_INNER]))}},
	B_TCPFOPS_T_WRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_TCPFOPS_T_WRITE]))}},
	B_SYS_GETTIMEOFDAY: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETTIMEOFDAY]))}},
	B_SYS_GETRUSAGE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETRUSAGE]))}},
	B_SYS_PWRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_PWRITE]))}},
	B_FS_T__FULLPATH: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_FS_T__FULLPATH]))}},
	B_IMEMNODE_T_IREAD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T_IREAD]))}},
	B_ARP_RESOLVE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_ARP_RESOLVE]))}},
	B_IMEMNODE_T_IMMAPINFO: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T_IMMAPINFO]))}},
	B_PIPE_T_OP_FDADD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_PIPE_T_OP_FDADD]))}},
	B_SYS_ACCESS: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_ACCESS]))}},
	B_SYS_LISTEN: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_LISTEN]))}},
	B_SYS_EXECV: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_EXECV]))}},
	B_SYS_GETRLIMIT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETRLIMIT]))}},
	B_SYS_SYNC: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SYNC]))}},
	B_FUTEX_T_FUTEX_START: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_FUTEX_T_FUTEX_START]))}},
	B_SYS_MKDIR: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_MKDIR]))}},
	B_SYS_KILL: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_KILL]))}},
	B_SYS_FUTEX: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_FUTEX]))}},
	B_USERIOVEC_T_IOV_INIT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_USERIOVEC_T_IOV_INIT]))}},
	B_SYS_OPEN: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_OPEN]))}},
	B_SYS_MMAP: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_MMAP]))}},
	B_SYS_DUP2: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_DUP2]))}},
	B_SYS_ACCEPT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_ACCEPT]))}},
	B_SYS_SHUTDOWN: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SHUTDOWN]))}},
	B_FS_T__ISANCESTOR: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_FS_T__ISANCESTOR]))}},
	B_PIPEFOPS_T_WRITE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_PIPEFOPS_T_WRITE]))}},
	B_SYS_GETPID: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETPID]))}},
	B_SYS_READ: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_READ]))}},
	B_SYS_THREXIT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_THREXIT]))}},
	B_FS_T_FS_RENAME: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_FS_T_FS_RENAME]))}},
	B_FS_T_FS_NAMEI: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_FS_T_FS_NAMEI]))}},
	B_IMEMNODE_T_BMAPFILL: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T_BMAPFILL]))}},
	B_SYS_SIGACTION: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SIGACTION]))}},
	B_SYS_GETPPID: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_GETPPID]))}},
	B_SYS_SETSOCKOPT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_SETSOCKOPT]))}},
	B_SYS_FORK: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_FORK]))}},
	B_SYS_WAIT4: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYS_WAIT4]))}},
	B_SYSCALL_T_SYS_CLOSE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYSCALL_T_SYS_CLOSE]))}},
	B_SYSCALL_T_SYS_EXIT: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_SYSCALL_T_SYS_EXIT]))}},
	B_IMEMNODE_T_IFREE: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IMEMNODE_T_IFREE]))}},
	B_BITMAP_T_APPLY: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_BITMAP_T_APPLY]))}},
	B_ICACHE_T_FREEDEAD: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_ICACHE_T_FREEDEAD]))}},
	B_IXGBE_T_INT_HANDLER: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_IXGBE_T_INT_HANDLER]))}},
	B_AHCI_PORT_T_QUEUEMGR: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_AHCI_PORT_T_QUEUEMGR]))}},
	B_AHCI_DISK_T_INT_HANDLER: &res.Res_t{Objs: runtime.Resobjs_t{1: uint32(uint(bounds[B_AHCI_DISK_T_INT_HANDLER]))}},
}

var bounds = []int{
	B_SYS_PAUSE:           0,
	B_SYS_PREAD:           96*24 + 1*4096 + 1*20 + 97*16 + 100*32 + 289*40 + 105*96 + 3*208 + 98*48 + 1*112 + 2*64 + 108*80,
	B_SYS_POLL:            (1024)*240 + (512)*32 + 66*32 + 1*4096 + 70*96 + 67*48 + 1*4120 + 193*40 + 72*80 + 2*208 + 2*64 + 65*24 + 65*16 + 1*20,
	B_SYS_WRITEV:          1*4120 + 1*4096 + 2*64 + 258*48 + 770*40 + 280*96 + 289*80 + 257*16 + 1*20 + 257*24 + 1*184 + 259*32 + 8*208,
	B_SYS_PROF:            1*48 + 64*1048 + 1*536 + 1*64 + 66*16 + 1*65536 + 2*32792,
	B_SYS_TRUNCATE:        6*208 + 469*16 + 1*8 + 1*20 + 2*1 + 283*48 + 314*80 + 280*32 + 308*96 + 751*64 + 280*24 + 834*40 + 1*4096,
	B_SYS_SOCKET:          1*16 + 1*576 + 1*24 + 1*4120 + 1*144 + 2*56,
	B_SYS_BIND:            214*24 + 238*96 + 2*104 + 2*8 + 403*16 + 1*56 + 1*1248 + 1*280 + 214*48 + 215*32 + 1*16408 + 1*11288 + 642*40 + 749*64 + 4*208 + 1*184 + 242*80,
	B_SYS_CHDIR:           738*40 + 2*1 + 751*64 + 251*48 + 278*80 + 273*96 + 1*8 + 1*4096 + 1*20 + 248*32 + 437*16 + 248*24 + 5*208,
	B_RAWDFOPS_T_WRITE:    (84)*40 + 1*40 + 1*4096 + 1*20 + 28*16 + 28*24 + 35*80 + 33*96 + 3*8 + 2*64 + 2*208 + 30*48 + 34*32,
	B_RAWDFOPS_T_READ:     (81)*40 + 27*16 + 35*80 + 2*208 + 1*40 + 1*4096 + 33*96 + 33*32 + 27*24 + 29*48 + 3*8 + 2*64 + 1*20,
	B_SYS_READV:           290*48 + 1*1024 + 4*25 + 324*80 + 288*24 + 9*208 + 2*64 + 1*4120 + 866*40 + 1*4096 + 1*184 + 1*20 + 289*16 + 315*96 + 291*32,
	B_SYS_SENDTO:          4*8 + 1*16408 + 243*80 + 218*24 + 4*104 + 220*16 + 654*40 + 220*32 + 187*64 + 7*208 + 236*96 + 1*11288 + 218*48 + 1*280,
	B_SYS_RECVFROM:        1*16408 + 290*32 + 290*16 + 288*48 + 864*40 + 4*8 + 9*208 + 315*96 + 1*11288 + 1*4120 + 4*104 + 288*24 + 324*80,
	B_SYS_SETRLIMIT:       193*40 + 72*80 + 66*48 + 70*96 + 1*4096 + 1*20 + 2*208 + 64*24 + 65*16 + 66*32 + 2*64,
	B_INSERTARGS:          64*1560 + 1*1 + 2*208 + 1*280 + 1*536 + 66*48 + 70*96 + 66*32 + 2*64 + 72*80 + 1*112 + 65*16 + 193*40 + 129*24 + 1*4096 + 1*20,
	B_SYS_WRITE:           140*96 + 128*48 + 1*16408 + 4*208 + 2*104 + 2*8 + 145*80 + 1*11288 + 130*16 + 129*24 + 129*32 + 384*40,
	B_SYS_GETCWD:          119*16 + 348*40 + 1*4096 + 3*1 + 132*80 + 119*32 + 118*48 + 4*208 + 128*96 + 377*64 + 1*8 + 117*24 + 1*20,
	B_SYS_SOCKETPAIR:      1*4096 + 2*4120 + 140*96 + 385*40 + 2*184 + 4*56 + 130*32 + 129*16 + 130*48 + 1*20 + 134*24 + 4*64 + 144*80 + 4*208,
	B_SYS_LINK:            751*64 + 366*96 + 338*24 + 1*20 + 1*8 + 2*1 + 1008*40 + 338*32 + 525*16 + 9*208 + 375*80 + 339*48 + 1*4096,
	B_SYS_MKNOD:           5*208 + 250*48 + 1*8 + 273*96 + 2*1 + 437*16 + 751*64 + 247*24 + 278*80 + 248*32 + 738*40 + 1*4096 + 1*20,
	B_SYS_GETTID:          0,
	B_SYS_UNLINK:          937*64 + 254*96 + 233*48 + 5*208 + 235*16 + 232*32 + 1*8 + 1*4096 + 1*20 + 259*80 + 690*40 + 232*24 + 2*1,
	B_SYS_REBOOT:          0,
	B_KBD_DAEMON:          1*16 + 1*1024 + 1*24 + 1*32 + 4*25 + 1*1,
	B_SYS_FSTAT:           1*4096 + 49*96 + 48*32 + 2*208 + 1*72 + 48*48 + 139*40 + 2*64 + 1*20 + 47*16 + 46*24 + 51*80,
	B_SYS_MUNMAP:          1*112 + 1*80 + 2*56 + 1*144,
	B_SYS_RECVMSG:         864*24 + 1*184 + 1*4096 + 1*4120 + 27*208 + 866*16 + 2*174 + 2*64 + 2594*40 + 945*96 + 972*80 + 866*48 + 869*32 + 1*20,
	B_SYS_FCNTL:           0,
	B_SYS_RENAME:          1809*16 + 975*96 + 1*20 + 874*32 + 12*208 + 987*80 + 5*8 + 2*1 + 3370*64 + 4*56 + 1*4096 + 873*48 + 873*24 + 2610*40,
	B_SYS_FTRUNCATE:       1*20 + 36*80 + 32*24 + 35*32 + 35*96 + 1*4096 + 2*64 + 33*16 + 96*40 + 34*48 + 1*208,
	B_SYS_NANOSLEEP:       140*96 + 4*208 + 1*4096 + 2*64 + 385*40 + 128*24 + 144*80 + 130*32 + 1*20 + 131*48 + 129*16,
	B_SYS_PIPE2:           1*184 + 1*4096 + 129*16 + 1*20 + 132*24 + 385*40 + 140*96 + 2*56 + 4*208 + 2*4120 + 130*48 + 144*80 + 130*32 + 2*64,
	B_SYS_STAT:            157*32 + 1*8 + 166*96 + 189*64 + 1*4096 + 1*20 + 5*208 + 2*1 + 171*80 + 1*72 + 462*40 + 156*48 + 157*16 + 156*24,
	B_SYS_LSEEK:           18*16 + 1*208 + 17*48 + 1*72 + 2*8 + 2*64 + 42*40 + 16*32 + 16*80 + 14*96 + 15*24 + 6*104 + 1*20,
	B_SYS_CONNECT:         90*48 + 99*80 + 187*64 + 91*32 + 2*56 + 92*16 + 2*104 + 1*280 + 1*11288 + 3*208 + 92*24 + 1*184 + 2*8 + 270*40 + 96*96 + 1*16408,
	B_SYS_SENDMSG:         831*96 + 1*280 + 855*80 + 1*88 + 2*1 + 764*24 + 1*184 + 24*208 + 767*32 + 1*8 + 1*4096 + 1*20 + 764*48 + 2287*40 + 189*64 + 1*72 + 765*16,
	B_SYS_GETSOCKOPT:      481*40 + 175*96 + 5*208 + 2*64 + 163*32 + 161*16 + 180*80 + 162*48 + 1*8 + 1*4096 + 1*20 + 160*24,
	B_SYS_INFO:            1*32 + 1*5776,
	B_SYS_GETTIMEOFDAY:    34*32 + 32*24 + 2*64 + 1*4096 + 1*20 + 98*40 + 33*16 + 36*80 + 1*208 + 34*48 + 35*96,
	B_SYS_GETRUSAGE:       33*16 + 35*96 + 33*24 + 36*80 + 1*56 + 34*48 + 1*4096 + 2*64 + 1*20 + 34*32 + 1*208 + 97*40,
	B_SYS_PWRITE:          2*64 + 295*40 + 101*48 + 108*96 + 99*16 + 3*208 + 1*4096 + 1*20 + 102*32 + 98*24 + 110*80,
	B_ARP_RESOLVE:         1*24 + 1*6 + 1*48 + 1*32 + 2*8 + 1*8216 + 1*42,
	B_SYS_ACCESS:          437*16 + 2*1 + 1*4096 + 738*40 + 751*64 + 5*208 + 273*96 + 278*80 + 250*48 + 1*8 + 1*20 + 248*32 + 247*24,
	B_SYS_LISTEN:          2*4120 + 1*136 + 1*56 + 1*75776,
	B_SYS_EXECV:           4*8 + 1*1 + 1*536 + 751*64 + 665*48 + 728*96 + 1*56 + 746*80 + 1*512 + 1989*40 + 64*1560 + 664*32 + 18*208 + 1*280 + 1*4096 + 735*24 + 1*1048 + 7*112 + 850*16 + 1*20,
	B_SYS_GETRLIMIT:       128*24 + 130*48 + 385*40 + 140*96 + 1*4096 + 2*64 + 1*20 + 129*16 + 144*80 + 130*32 + 4*208,
	B_SYS_SYNC:            0,
	B_SYS_MKDIR:           369*32 + 5*208 + 1098*40 + 369*48 + 1*8 + 2*1 + 416*80 + 1314*64 + 928*16 + 367*24 + 411*96 + 1*4096 + 1*20,
	B_SYS_KILL:            0,
	B_SYS_FUTEX:           144*80 + 129*16 + 1*4096 + 12*104 + 133*48 + 2*64 + 3*8 + 388*40 + 1*232 + 12*424 + 1*20 + 4*208 + 136*24 + 130*32 + 140*96,
	B_SYS_OPEN:            437*16 + 278*80 + 5*208 + 273*96 + 2*1 + 1*4096 + 248*32 + 751*64 + 738*40 + 251*48 + 248*24 + 1*4120 + 1*8 + 1*20,
	B_SYS_MMAP:            1*24 + 2*112 + 1*80 + 2*56 + 1*144,
	B_SYS_DUP2:            1*24 + 2*56 + 1*144,
	B_SYS_ACCEPT:          161*32 + 5*208 + 163*16 + 2*8 + 1*11288 + 1*16408 + 180*80 + 175*96 + 1*4120 + 2*104 + 160*48 + 480*40 + 161*24,
	B_SYS_SHUTDOWN:        2*56 + 1*144,
	B_SYS_GETPID:          0,
	B_SYS_READ:            162*16 + 2*104 + 480*40 + 161*32 + 1*4120 + 1*11288 + 1*16408 + 5*208 + 175*96 + 160*24 + 180*80 + 160*48 + 2*8,
	B_SYS_THREXIT:         1*24 + 1*8 + 2*56 + 1*144,
	B_SYS_SIGACTION:       0,
	B_SYS_GETPPID:         0,
	B_SYS_SETSOCKOPT:      1*20 + 64*24 + 2*208 + 65*16 + 72*80 + 66*48 + 67*32 + 70*96 + 193*40 + 1*4096 + 2*64,
	B_SYS_FORK:            512*2*24 + 512*2*112 + 512*24 + 1*208 + 36*80 + 1*4096 + 3*56 + 1*360 + 34*32 + 2*64 + 97*40 + 35*96 + 1*192 + 1*16384 + 2*8 + 1*20 + 34*48 + 32*24 + 3*1 + 33*16 + 1*4120,
	B_SYS_WAIT4:           96*24 + 99*48 + 105*96 + 1*56 + 1*4096 + 2*64 + 98*32 + 3*208 + 97*16 + 289*40 + 108*80 + 1*20,
	B_SYSCALL_T_SYS_CLOSE: 2*56 + 1*144,
	B_SYSCALL_T_SYS_EXIT:  1*24 + 1*8 + 2*56 + 1*144,

	// deeper checks
	B_PROC_T_USERARGS:               (117)*40 + 1*4096 + 1*20 + 3*208 + 9*1 + 40*16 + 1*40 + 1*1048 + 48*96 + 51*80 + 54*32 + 41*48 + 2*64 + 40*24 + 7*8,
	B_USERIOVEC_T__TX:               (78)*40 + 32*96 + 26*16 + 1*40 + 1*4096 + 26*24 + 28*48 + 3*8 + 32*32 + 34*80 + 2*208 + 2*64 + 1*20,
	B_IMEMNODE_T_IFREE:              (39)*40 + 13*24 + 13*96 + 13*80 + 15*16 + 2*64 + 15*48 + 1*90 + 16*32 + 1*20,
	B_PIPE_T_OP_FDADD:               1 * 80,
	B_PIPEFOPS_T_WRITE:              (156)*40 + 64*96 + 68*80 + 4*208 + 5*8 + 1*40 + 2*64 + 60*32 + 54*48 + 52*24 + 52*16 + 1*4096 + 1*20,
	B_IMEMNODE_T__DESCAN:            (36)*40 + 15*80 + 12*16 + 14*48 + 13*24 + 15*32 + 1*4096 + 1*8 + 15*96 + 189*64 + 1*20,
	B_IMEMNODE_T_DO_WRITE:           (117)*40 + 2*208 + 39*16 + 1*4096 + 1*20 + 50*80 + 42*48 + 3*8 + 1*40 + 2*64 + 39*24 + 45*32 + 49*96,
	B_FS_T__ISANCESTOR:              (42)*40 + 16*16 + 15*24 + 4*8 + 1*4096 + 16*48 + 189*64 + 19*80 + 2*208 + 17*96 + 2*1 + 1*20 + 19*32,
	B_IMEMNODE_T_IREAD:              (114)*40 + 1*40 + 3*8 + 1*112 + 38*16 + 2*208 + 40*48 + 49*80 + 44*32 + 1*20 + 47*96 + 38*24 + 1*4096 + 2*64,
	B_IMEMNODE_T_IMMAPINFO:          (36)*40 + 15*32 + 1*20 + 15*80 + 12*24 + 15*96 + 14*48 + 1*4096 + 1*8 + 2*64 + 1*40 + 12*16,
	B_TCPFOPS_T_WRITE:               (156)*40 + 52*16 + 68*80 + 52*24 + 60*32 + 64*96 + 1*20 + 5*8 + 4*208 + 54*48 + 1*40 + 1*4096 + 2*64,
	B_FS_T_FS_NAMEI:                 (42)*40 + 16*16 + 16*24 + 189*64 + 19*32 + 2*1 + 1*4096 + 1*20 + 4*8 + 19*80 + 17*96 + 16*48 + 2*208,
	B_IMEMNODE_T_IWRITE:             (114)*40 + 38*16 + 44*32 + 1*40 + 2*64 + 38*24 + 40*48 + 1*4096 + 1*20 + 2*208 + 48*96 + 49*80 + 3*8,
	B_TCPFOPS_T_READ:                (156)*40 + 60*32 + 5*8 + 68*80 + 4*208 + 1*40 + 1*4096 + 2*64 + 54*48 + 52*24 + 52*16 + 64*96 + 1*20,
	B_FUTEX_T_FUTEX_START:           (1560)*40 + 1*64 + 560*32 + 41*8 + 1*40 + 521*48 + 680*80 + 4*424 + 571*24 + 20*112 + 4*104 + 40*208 + 520*16 + 2*232 + 640*96,
	B_USERBUF_T__TX:                 (39)*40 + 17*80 + 2*8 + 17*32 + 13*16 + 16*96 + 1*208 + 1*20 + 13*24 + 15*48 + 1*40 + 1*4096 + 2*64,
	B_FS_T_FS_RENAME:                (1161)*40 + 1*4096 + 1*20 + 480*96 + 400*32 + 488*80 + 3370*64 + 2*1 + 14*8 + 1325*16 + 4*56 + 390*48 + 8*208 + 390*24,
	B_TCPTIMERS_T__TCPTIMERS_DAEMON: 1*48 + 360*8 + 360*32 + 1*24 + 2*56 + 1*144,
	B_ELF_T_ELF_LOAD:                (156)*40 + 68*80 + 59*32 + 52*24 + 52*16 + 1*20 + 2*64 + 5*8 + 5*40 + 55*48 + 4*208 + 1*504 + 64*96 + 4*112 + 1*4096,
	B_PROC_T_USER2K_INNER:           (39)*40 + 15*48 + 17*80 + 13*16 + 16*96 + 1*208 + 13*24 + 1*20 + 17*32 + 2*8 + 1*40 + 1*4096 + 2*64,
	B_IMEMNODE_T_BMAPFILL:           (33)*40 + 11*24 + 13*48 + 1*8 + 2*64 + 1*20 + 14*96 + 14*80 + 11*16 + 14*32 + 1*4096,
	B_FS_T__FULLPATH:                (78)*40 + 34*80 + 28*16 + 27*24 + 32*96 + 1*20 + 28*48 + 3*8 + 2*208 + 30*32 + 377*64 + 2*1 + 1*4096,
	B_IXGBE_T_INT_HANDLER:           2*4120 + 1*16384 + 256*32 + 514*8 + 256*8216 + 257*48 + 1*16,
	B_PROC_T_K2USER_INNER:           (39)*40 + 1*40 + 1*20 + 1*208 + 2*8 + 16*96 + 13*24 + 1*4096 + 2*64 + 17*80 + 17*32 + 15*48 + 13*16,
	B_USERIOVEC_T_IOV_INIT:          (156)*40 + 64*96 + 1*40 + 1*4096 + 2*64 + 54*48 + 1*184 + 59*32 + 52*16 + 5*8 + 52*24 + 68*80 + 1*20 + 4*208,

	// debug reservations
	B_AHCI_PORT_T_QUEUEMGR:    1*20 + 1*48 + 1*32 + 1*64,
	B_AHCI_DISK_T_INT_HANDLER: 96*16 + 1*32 + 1*48,

	// cache/forced eviction allocations
	B_ICACHE_T_FREEDEAD: (39)*40 + 16*16 + 13*96 + 1*90 + 13*80 + 15*48 + 13*24 + 16*32 + 2*64 + 1*20,
	B_BITMAP_T_APPLY:    (3)*40 + 2*64 + 1*20 + 1*16 + 3*48 + 1*24 + 2*32 + 1*80 + 1*96,
}
