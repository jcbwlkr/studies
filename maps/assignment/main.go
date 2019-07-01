package main

import "fmt"

type game struct {
	started bool
}

var games = make(map[string]game)

func main() {

	games["jcb"] = game{started: true}

	games["jcb"].started = false

	fmt.Println(games["jcb"].started)
}
