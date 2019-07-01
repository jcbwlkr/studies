package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}

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

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {

			It("should be a novel", func() {
				fmt.Println("Shortbook title is", shortBook.Title)
				shortBook.Title = "Cat in the Hat"
				fmt.Println("Shortbook title is changed to", shortBook.Title)

				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				fmt.Println("Shortbook title is returned to", shortBook.Title)
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})
	})
})
