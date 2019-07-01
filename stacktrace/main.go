package main

type Foo struct{}

func (f *Foo) Bar(name string) {
	panic("omg")
}

func slice(b []byte) {
	panic("die")
}

func main() {
	//var f Foo
	//f.Bar("hi")
	s := []byte("abcd")
	slice(s)
}
