package main

import "testing"

type Book struct {
	Title, Author string
	Pages         int
}

func (b Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}

func shortBook() Book {
	return Book{
		Title:  "Fox In Socks",
		Author: "Dr. Seuss",
		Pages:  24,
	}
}

func longBook() Book {
	return Book{
		Title:  "Les Miserables",
		Author: "Victor Hugo",
		Pages:  1488,
	}
}

func TestBookLength(t *testing.T) {
	tests := []struct {
		name string
		book Book
		want string
	}{
		{"short book", shortBook(), "SHORT STORY"},
		{"long book", longBook(), "NOVEL"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.book.CategoryByLength(); got != test.want {
				t.Errorf("CategoryByLength() for less than 300 pages should be %q, got %q", test.want, got)
			}
		})
	}
}
