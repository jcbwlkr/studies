package channels

import (
	"errors"
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	c := make(chan error, 1)

	Foo("work", c)

	close(c)

	err := <-c
	if err != nil {
		t.Errorf("there should be no error on this channel but there was %v", err)
	}
}

func TestFooBad(t *testing.T) {
	c := make(chan error, 1)

	Foo("break", c)

	close(c)

	err := <-c
	if err == nil {
		t.Error("there should be an error in the channel but there was not")
	}

}

func Foo(task string, c chan error) {
	if task == "break" {
		fmt.Println("this is going to break")
		c <- errors.New("this is broken")
		return
	}

	fmt.Println("this works")
}
