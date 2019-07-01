package listtimes

import (
	"strconv"
	"testing"
)

type userA struct {
	name string
}

type userB struct {
	name string
}

func getUsers() []userA {
	users := make([]userA, 100000)
	for i := 0; i < 100000; i++ {
		users[i] = userA{name: "foo" + strconv.Itoa(i)}
	}

	return users
}

func BenchmarkAppend(b *testing.B) {
	users := getUsers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := []userB{}
		for _, u := range users {
			c = append(c, userB{name: u.name})
		}
	}
}

func BenchmarkMakeLength(b *testing.B) {
	users := getUsers()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := make([]userB, len(users))
		for i, u := range users {
			c[i] = userB{name: u.name}
		}
	}
}
