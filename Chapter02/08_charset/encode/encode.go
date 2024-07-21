package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	f, err := os.OpenFile("out.txt", os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Decode to unicode
	encoder := charmap.Windows1250.NewEncoder()
	writer := encoder.Writer(f)
	_, err = io.WriteString(writer, "Gdańsk")
	if err != nil {
		log.Fatal(err)
	}
}
