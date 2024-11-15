package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	check(cmd.Run())

	file, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Print(string(file))
}
