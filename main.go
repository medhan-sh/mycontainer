package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Runningggg %v at PID: %v\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Run()

}
func child() {
	fmt.Printf("Child func: setting namespace\n PID : %v\n", os.Getpid())
	if err := syscall.Sethostname([]byte("inside")); err != nil {
		fmt.Println("Error setting namespace : ", err)
	}
	if err := syscall.Chroot("/container"); err != nil {
		fmt.Println("Error changing root : ", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		fmt.Println("Error changing directory : ", err)
	}
	if err := syscall.Mount("proc", "proc", "proc", 0, ""); err != nil {
		fmt.Println("Error mounting : ", err)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/bash command - %v\n", err)
		os.Exit(1)
	}
	defer syscall.Unmount("proc", 0)

}
