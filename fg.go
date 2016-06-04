package main

import (
	"fmt"
	//"log"
	"os"
	"os/exec"
	"syscall"
)

func ioctl(fd, call, data uintptr) error {
	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, call, data)
	err := error(nil)
	if errno != 0 {
		err = errno
	}
	err = os.NewSyscallError("SYS_IOCTL", err)
	if r1 != 0 {
		return fmt.Errorf("r1 == %d", r1)
	}
	return err
}

func main() {
	//tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0777)
	//if err != nil {
	//	log.Fatalf("Can't test Foreground. Couldn't open /dev/tty: %s", err)
	//}
	cmd := &exec.Cmd{
		Path:   "/bin/vim",
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmd.Run()
}
