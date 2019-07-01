package main

import "fmt"

type fooer interface {
	foo()
}

type bar struct {
	calls int
}

func (b *bar) foo() {
	b.calls++
	fmt.Println("called bar.foo")
}

func main() {
	b := &bar{}

	do(b)

	fmt.Println(b.calls)

	var f fooer = b
	doPtr(&f)

	fmt.Println(b.calls)
}

func do(f fooer) {
	f.foo()
}

func doPtr(f *fooer) {
	(*f).foo()
}
