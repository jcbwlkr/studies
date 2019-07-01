package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {

	f, err := os.Create("books.json")
	if err != nil {
		log.Fatalln("could not make output file", err)
	}

	data := []Book{
		{Title: "Book 1", Author: "Author 1"},
		{Title: "Book 2", Author: "Author 1"},
		{Title: "Book 3", Author: "Author 2"},
		{Title: "Book 4", Author: "Author 2"},
	}

	if err := json.NewEncoder(f).Encode(data); err != nil {
		log.Fatalln("could not write to output file", err)
	}

	fmt.Println("Done! Books written to books.json")
}
