package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type thing struct {
	Name string
	Data []byte
}

func main() {
	t := thing{Name: "Test", Data: []byte(`foobar`)}

	fmt.Printf("Thing: %+v\n", t)

	s, _ := json.Marshal(t)

	fmt.Println("JSON: " + string(s))

	var m map[string]interface{}

	json.Unmarshal(s, &m)

	fmt.Printf("Map: %+v\n", m)

	want := map[string]interface{}{
		"Data": "Zm9vYmFy",
		"Name": "Test",
	}

	fmt.Println(reflect.DeepEqual(want, m))
}
