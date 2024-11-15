package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := exec.Command("sh", "-c", `echo "$HOME"`)
	out, _ := c.CombinedOutput()

	fmt.Println(string(out))
}
