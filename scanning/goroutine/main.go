package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Inputscanner(c chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scanner died:", err)
	}
	return
}

func main() {

	c := make(chan string)

	stop := "false"

	go Inputscanner(c)

RESET:
	for stop == "false" {
		fmt.Println("Test")
		time.Sleep(1 * time.Second)
		select {
		case msg := <-c:
			fmt.Printf(msg)
			if msg == "true" {
				fmt.Println("DONE")
				stop = "true"
			}
		default:
			continue RESET
		}
	}
}
