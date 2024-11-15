package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"regexp"
)

func main() {
	input, err := os.Open("init.md")
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	out := csv.NewWriter(output)

	re := regexp.MustCompile(`^- \[([^\]]+)\]\(([^)]+)`)

	s := bufio.NewScanner(input)
	for s.Scan() {
		ln := s.Text()

		if !re.MatchString(ln) {
			continue
		}
		matches := re.FindAllStringSubmatch(ln, -1)

		row := []string{
			matches[0][1],
			"",
			matches[0][2],
		}
		if err := out.Write(row); err != nil {
			log.Print(err)
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalln("scan error", err)
	}

	out.Flush()
	if err := out.Error(); err != nil {
		log.Fatalln("write error", err)
	}
}
