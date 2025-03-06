package main

import "fmt"

type Customer struct {
	ID   int
	Name string
}

type Customers []Customer

func (customers *Customers) Add(customer Customer) {

	for i, c := range *customers {
		if c.ID == customer.ID {
			(*customers)[i] = customer
			return
		}
	}

	*customers = append(*customers, customer)

}

func main() {

	var c Customers

	fmt.Printf("Addr: %p Len: %d Cap: %d\n", &c, len(c), cap(c))

	c.Add(Customer{ID: 1, Name: "Jacob"})
	fmt.Println(c[0])
	c.Add(Customer{ID: 1, Name: "Jake"})
	fmt.Println(c[0])

	c.Add(Customer{ID: 2, Name: "Jenn"})
	c.Add(Customer{ID: 3, Name: "Kell"})
	c.Add(Customer{ID: 4, Name: "Carter"})
	c.Add(Customer{ID: 5, Name: "Rory"})

	fmt.Printf("Addr: %p Len: %d Cap: %d\n", &c, len(c), cap(c))
}
