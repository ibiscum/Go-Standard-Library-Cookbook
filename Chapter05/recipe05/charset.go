package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	// Write the string
	// encoded to Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String("This is sample text with runes Š")
	if e != nil {
		panic(e)
	}

	e = os.WriteFile("example.txt", []byte(s), os.ModePerm)
	if e != nil {
		panic(e)
	}

	// Decode to UTF-8
	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()

	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
