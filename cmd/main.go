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
		panic("Invalid command input")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	// Remap IO
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Add Linux Namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	cmd.Run()
}

func child() {
	fmt.Printf("Sticky running %v with pid: %d\n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("container"))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	// Remap IO
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
