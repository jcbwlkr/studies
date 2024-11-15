package main

import "fmt"

func printResult(msg string, history [][]int) {
	fmt.Println(msg)

	size := len(history)

	fmt.Print("╔═════╦")
	for i := 0; i < size; i++ {
		if i > 0 {
			fmt.Print("═")
		}
		fmt.Print("═════")
	}
	fmt.Println("╗")

	// Rows with Y labels, in descending order
	for y := size - 1; y >= 0; y-- {
		fmt.Printf("║ % 3d ║", y)
		for x := 0; x < size; x++ {
			if x > 0 {
				fmt.Print("┊")
			}
			fmt.Printf(" % 3d ", history[x][y])
		}
		fmt.Println("║")

		if y > 0 {
			fmt.Print("║┈┈┈┈┈║")
			for x := 0; x < size; x++ {
				if x > 0 {
					fmt.Print("┊")
				}
				fmt.Print("┈┈┈┈┈")
			}
			fmt.Println("║")
		}
	}

	fmt.Print("╠═════╬")
	for i := 0; i < size; i++ {
		if i > 0 {
			fmt.Print("═")
		}
		fmt.Print("═════")
	}
	fmt.Println("╣")

	// Column labels
	fmt.Print("║  ♞  ║")
	for x := 0; x < size; x++ {
		if x > 0 {
			fmt.Print("┊")
		}
		fmt.Printf(" % 3d ", x)
	}
	fmt.Println("║")

	fmt.Print("╚═════╩")
	for i := 0; i < size; i++ {
		if i > 0 {
			fmt.Print("═")
		}
		fmt.Print("═════")
	}
	fmt.Println("╝")
}

/*
╔═╤═╤═╤═╤═╤═╤═╤═╗╮
║♜│♞│♝│♛│♚│♝│♞│♜║8
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║♟│♟│♟│♟│♟│♟│♟│♟║7
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║ │░│ │░│ │░│ │░║6
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║░│ │░│ │░│ │░│ ║5
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║ │░│ │░│ │░│ │░║4
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║░│ │░│ │░│ │░│ ║3
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║♙│♙│♙│♙│♙│♙│♙│♙║2
╟─┼─┼─┼─┼─┼─┼─┼─╢┊
║♖│♘│♗│♕│♔│♗│♘│♖║1
╚═╧═╧═╧═╧═╧═╧═╧═╝┊
╰a┈b┈c┈d┈e┈f┈g┈h┈╯
*/
