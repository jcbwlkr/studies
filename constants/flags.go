package main

import "fmt"

// Example constants
const (
	Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
	Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
	Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
	Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
	Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
	LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
)

func main() {
	fmt.Println("vim-go")
}
