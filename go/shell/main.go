package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	shell := exec.Command("/bin/bash")
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Run()
	fmt.Println("Ran a shell in go")

}
