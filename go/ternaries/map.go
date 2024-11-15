package main

func main() {
	color := 80
	str := map[bool]string{true: "red", false: "blue"}[color > 100]
	println(str)
}
