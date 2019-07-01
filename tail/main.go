package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("must provide path to a file")
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
	}

	fmt.Println("vim-go")
}
