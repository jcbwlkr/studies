package main

import "fmt"

// dataSource describes the info about a subcription
type dataSource struct {
	Email string
}

// subcriptions is a slice of all the current subcriptions.
type subcriptions []dataSource

func (s *subcriptions) Remove(i int) {
	t := *s                        // get the value of s
	t = append(t[0:i], t[i+1:]...) // shrink it
	*s = t                         // replace the value of s
}

func main() {

	s := subcriptions{
		{Email: "foo"},
		{Email: "bar"},
		{Email: "baz"},
		{Email: "qux"},
		{Email: "abc"},
		{Email: "def"},
	}

	inspect(s)
	s.Remove(2)
	inspect(s)
}

func inspect(s subcriptions) {
	fmt.Printf("len[%d], cap[%d], vals: %+v\n", len(s), cap(s), s)
}
