package main

import (
	"fmt"
	"reflect"
)

type BlogPost struct {
	Title    string
	Body     string
	Comments map[string]string
}

func main() {
	a := BlogPost{
		Title: "Awesome Thing",
		Body:  "Go is cool!",
		Comments: map[string]string{
			"joe-blow":          "I dunno, it doesn't have generics!",
			"frank-fart":        "Needs more enterprise",
			"hater-mc-hateface": "I just like things my way",
		},
	}

	b := BlogPost{
		Body: "Go is cool!",
		Comments: map[string]string{
			"frank-fart":        "Needs more enterprise",
			"hater-mc-hateface": "I just like things my way",
			"joe-blow":          "I dunno, it doesn't have generics!",
		},
		Title: "Awesome Thing",
	}

	fmt.Println(reflect.DeepEqual(a, b))
}
