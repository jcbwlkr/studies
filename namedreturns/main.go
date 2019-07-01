package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Return value 1:", work(true))
	fmt.Println("Return value 2:", work(false))
}

func work(succeed bool) (err error) {

	n, err := otherWork()
	if err != nil {
		return err
	}
	_ = n

	defer func() {
		fmt.Println("In defer and err is", err)
		// LOL JK
		//err = nil
	}()

	if !succeed {
		err := errors.New("ERROR (SHADOWED)")
		return err
	}

	return nil
}

func otherWork() (int, error) {
	return 0, nil
}
