package proc

import "defs"
import "fdops"

type Syscall_i interface {
	Syscall(p *Proc_t, tid defs.Tid_t, tf *[defs.TFSIZE]uintptr) int
	Sys_close(proc *Proc_t, fdn int) int
	Sys_exit(Proc *Proc_t, tid defs.Tid_t, status int)
}

type Cons_i interface {
	Cons_read(ub fdops.Userio_i, offset int) (int, defs.Err_t)
	Cons_write(src fdops.Userio_i, off int) (int, defs.Err_t)
}
