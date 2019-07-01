package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	help := flag.Bool("h", false, "Print help and exit")
	helpLong := flag.Bool("help", false, "Print help and exit")

	host := flag.String("host", os.Getenv("HOST"), "An ip:port value for the web service")

	var cfg struct {
		Host string
	}

	flag.Parse()

	if *help || *helpLong {
		flag.Usage()
		os.Exit(0)
	}

	fmt.Println(*host)
}
