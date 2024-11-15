package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "test.sh")

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("Could not get stdout pipe", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalln("Could not start cmd", err)
	}

	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatalln("Could not create results file", err)
	}

	if _, err := io.Copy(f, pipe); err != nil {
		log.Fatalln("Could not copy pipe to file", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalln("command failed", err)
	}
}
