package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}

func main() {

	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)

	// Read the book one by one
	books := make([]Book, 0)
	for {
		tok, err := decoder.Token()
		if err != io.EOF && err != nil {
			log.Fatal("ERROR: ", err)
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "book" {
				// Decode the element to struct
				var b Book
				err = decoder.DecodeElement(&b, &tp)
				if err != nil {
					log.Fatal(err)
				}
				books = append(books, b)
			}
		}
	}
	fmt.Println(books)
}
