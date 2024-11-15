package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	src, dst := "foo", "newFoo"

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		newPath := strings.Replace(path, src, dst, 1)
		fmt.Println(path, "->", newPath)

		// If there was an error getting to the current path we can decide what to do about it
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.MkdirAll(newPath, info.Mode())
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		newF, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(newF, f)
		return err
	})

	if err != nil {
		log.Fatalln("could not walk whole tree", err)
	}
}
