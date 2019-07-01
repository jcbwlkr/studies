package main

import "fmt"

func main() {
	var errs []error
	inspect(errs)

	var x []error
	errs = append(errs, x...)
	errs = append(errs, []error{}...)

	inspect(errs)

	y := []error{
		nil,
		nil,
		nil,
	}
	errs = append(errs, y...)
	inspect(errs)
}

func inspect(e []error) {
	fmt.Printf("len[%d] cap[%d] vals: %v\n", len(e), cap(e), e)
}
