package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/flosch/pongo2"
)

func init() {
	pongo2.RegisterFilter("scream", Scream)
}

// Scream is a silly example of a filter function that upper cases strings
func Scream(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	if !in.IsString() {
		return nil, &pongo2.Error{
			ErrorMsg: "only strings should be sent to the scream filter",
		}
	}

	s := in.String()
	s = strings.ToUpper(in.String())

	return pongo2.AsValue(s), nil
}

func main() {
	tpl, err := pongo2.FromString("Hello {{ name|scream }}!")
	if err != nil {
		log.Fatal(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(pongo2.Context{"name": "stack overflow"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out) // Output: Hello STACK OVERFLOW!
}
