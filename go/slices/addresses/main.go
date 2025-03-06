package main

import (
	"fmt"
	"slices"
	"time"
)

type user struct {
	ID   int
	Date time.Time
}

func main() {

	users := []user{
		{ID: 1},
		{ID: 3},
		{ID: 5},
	}

	for i := 0; i < 7; i++ {

		idx := slices.IndexFunc(users, func(u user) bool {
			return u.ID == i
		})

		if idx < 0 {
			continue
		}

		u := &users[idx]
		u.Date = time.Now()
	}

	fmt.Printf("%+v\n", users)
}
