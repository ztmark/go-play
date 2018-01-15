package main

import (
    "os"
    "fmt"
    "os/exec"
    "syscall"
)

func main() {
    //switch os.Args[1] {
    //case "run":
    //    run()
    //default:
    //    fmt.Println("What?")
    //}

    syscall.Setenv("mark", "mark1")

}

func run() {
    fmt.Printf("Running %v\n", os.Args[2:])

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.SysProcAttr = &syscall.SysProcAttr{
        //Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
    }

    must(cmd.Start())

}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
