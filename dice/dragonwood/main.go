package main

import "fmt"

var sides = [...]int{1, 2, 2, 3, 3, 4}

var verbose = false

type result struct {
	min, max, total, rolls int
	mean                   float64
}

var results = make([]result, 6)

func main() {
	roll1()
	roll2()
	roll3()
	roll4()
	roll5()
	roll6()

	fmt.Println("+------------------------------------------------+")
	fmt.Println("| Dice |    1 |    2 |    3 |    4 |    5 |    6 |")
	fmt.Println("+------------------------------------------------+")
	fmt.Printf("| Min  | %4d | %4d | %4d | %4d | %4d | %4d |\n", results[0].min, results[1].min, results[2].min, results[3].min, results[4].min, results[5].min)
	fmt.Printf("| Mean |  %3.1f |  %3.1f |  %3.1f | %3.1f | %3.1f | %3.1f |\n", results[0].mean, results[1].mean, results[2].mean, results[3].mean, results[4].mean, results[5].mean)
	fmt.Printf("| Max  | %4d | %4d | %4d | %4d | %4d | %4d |\n", results[0].max, results[1].max, results[2].max, results[3].max, results[4].max, results[5].max)
	fmt.Println("+------------------------------------------------+")
}

func roll1() {
	if verbose {
		fmt.Println("Rolling 1 die")
	}
	var r result
	for _, d1 := range sides {
		if verbose {
			fmt.Println(d1)
		}

		r.rolls++
		r.total += d1

		if d1 > r.max {
			r.max = d1
		}
		if d1 < r.min || r.min == 0 {
			r.min = d1
		}
	}
	r.mean = float64(r.total) / float64(r.rolls)
	results[0] = r
}

func roll2() {
	if verbose {
		fmt.Println("Rolling 2 dice")
	}
	var r result
	for _, d1 := range sides {
		for _, d2 := range sides {
			val := d1 + d2
			if verbose {
				fmt.Println(d1, "+", d2, "=", val)
			}

			r.rolls++
			r.total += val

			if val > r.max {
				r.max = val
			}
			if val < r.min || r.min == 0 {
				r.min = val
			}
		}
	}
	r.mean = float64(r.total) / float64(r.rolls)
	results[1] = r
}

func roll3() {
	if verbose {
		fmt.Println("Rolling 3 dice")
	}
	var r result
	for _, d1 := range sides {
		for _, d2 := range sides {
			for _, d3 := range sides {
				val := d1 + d2 + d3
				if verbose {
					fmt.Println(d1, "+", d2, "+", d3, "=", val)
				}

				r.rolls++
				r.total += val

				if val > r.max {
					r.max = val
				}
				if val < r.min || r.min == 0 {
					r.min = val
				}
			}
		}
	}
	r.mean = float64(r.total) / float64(r.rolls)
	results[2] = r
}

func roll4() {
	if verbose {
		fmt.Println("Rolling 4 dice")
	}
	var r result
	for _, d1 := range sides {
		for _, d2 := range sides {
			for _, d3 := range sides {
				for _, d4 := range sides {
					val := d1 + d2 + d3 + d4
					if verbose {
						fmt.Println(d1, "+", d2, "+", d3, "+", d4, "=", val)
					}

					r.rolls++
					r.total += val

					if val > r.max {
						r.max = val
					}
					if val < r.min || r.min == 0 {
						r.min = val
					}
				}
			}
		}
	}
	r.mean = float64(r.total) / float64(r.rolls)
	results[3] = r
}

func roll5() {
	if verbose {
		fmt.Println("Rolling 5 dice")
	}
	var r result
	for _, d1 := range sides {
		for _, d2 := range sides {
			for _, d3 := range sides {
				for _, d4 := range sides {
					for _, d5 := range sides {
						val := d1 + d2 + d3 + d4 + d5
						if verbose {
							fmt.Println(d1, "+", d2, "+", d3, "+", d4, "+", d5, "=", val)
						}

						r.rolls++
						r.total += val

						if val > r.max {
							r.max = val
						}
						if val < r.min || r.min == 0 {
							r.min = val
						}
					}
				}
			}
		}
	}
	r.mean = float64(r.total) / float64(r.rolls)
	results[4] = r
}

func roll6() {
	if verbose {
		fmt.Println("Rolling 6 dice")
	}
	var r result
	for _, d1 := range sides {
		for _, d2 := range sides {
			for _, d3 := range sides {
				for _, d4 := range sides {
					for _, d5 := range sides {
						for _, d6 := range sides {
							val := d1 + d2 + d3 + d4 + d5 + d6
							if verbose {
								fmt.Println(d1, "+", d2, "+", d3, "+", d4, "+", d5, "+", d6, "=", val)
							}

							r.rolls++
							r.total += val

							if val > r.max {
								r.max = val
							}
							if val < r.min || r.min == 0 {
								r.min = val
							}
						}
					}
				}
			}
		}
	}
	r.mean = float64(r.total) / float64(r.rolls)
	results[5] = r
}
