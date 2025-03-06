package main

import (
	_ "embed" // for go:embed
	"fmt"

	"github.com/k3a/html2text"
)

//go:embed body.txt
var body string

func main() {

	fmt.Println(html2text.HTML2TextWithOptions(body, html2text.WithLinksInnerText()))
}
