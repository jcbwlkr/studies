package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var t = make(map[string]*template.Template)

func main() {

	t["vday"] = template.Must(template.ParseFiles("base.tmpl", "vday.tmpl"))
	t["xmas"] = template.Must(template.ParseFiles("base.tmpl", "xmas.tmpl"))

	run("vday")
	run("xmas")
}

func run(name string) {

	tm := t[name]
	if tm == nil {
		log.Print("Could not find template ", name)
		return
	}

	fmt.Println(name)

	fmt.Print("Subject: ")
	err := tm.ExecuteTemplate(os.Stdout, "subject", nil)
	fmt.Println("\nErr:", err)

	fmt.Print("Text: ")
	err = tm.ExecuteTemplate(os.Stdout, "text", nil)
	fmt.Println("\nErr:", err)

	fmt.Println("HTML:")
	err = tm.ExecuteTemplate(os.Stdout, "html-base", nil)
	fmt.Println("\nErr:", err)

	//fmt.Print("Body: ")
	//err = tm.ExecuteTemplate(os.Stdout, "body", nil)
	//fmt.Println("\nErr:", err)
}
