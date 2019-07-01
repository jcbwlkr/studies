package main

import "time"

var scores = make(map[string]int)

func main() {

	go player("Jacob")
	go player("Anna")
	go player("Kell")

	time.Sleep(5 * time.Second)
}

func player(name string) {
	for i := 0; i < 10000; i++ {
		scores[name]++
	}
}
