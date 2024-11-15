package main

import (
	"encoding/json"
	"fmt"
	"log"

	"go.opencensus.io/trace/propagation"
)

var payload struct {
	Data []byte
}

func main() {

	data := []byte(`{"Data": "AAChn6ZDOPPxmBJFeznL70oPAbD45oLAhR4mAgE="}`)

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Fatal(err)
	}

	sc, ok := propagation.FromBinary(payload.Data)
	if !ok {
		log.Fatal("Was not ok")
	}

	fmt.Printf("%+v\n", sc)
}
