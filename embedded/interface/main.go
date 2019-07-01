package main

type printer interface {
	Print()
}

type concrete struct {
	*printer
}

func main() {
}
