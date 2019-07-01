package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "env")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
