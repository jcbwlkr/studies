package main

import (
	"fmt"
)

func main() {

	monthdays := []struct {
		Name string
		Days int
	}{
		{Name: "Jan", Days: 31},
		{Name: "Feb", Days: 28},
		{Name: "March", Days: 31},
		{Name: "Apr", Days: 30},
		{Name: "May", Days: 31},
		{Name: "Jun", Days: 30},
		{Name: "Jul", Days: 31},
		{Name: "Aug", Days: 31},
		{Name: "Sep", Days: 30},
		{Name: "Oct", Days: 31},
		{Name: "Nov", Days: 30},
		{Name: "Dec", Days: 31},
	}

	var year int
	for _, m := range monthdays {
		fmt.Printf("%s has %d days\n", m.Name, m.Days)
		year += m.Days
	}

	fmt.Printf("Numbers of days in a year: %d\n", year)
}
