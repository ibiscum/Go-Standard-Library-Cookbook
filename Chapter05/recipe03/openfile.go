package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("temp/file.txt")
	if err != nil {
		log.Panic(err)
	}

	c, err := io.ReadAll(f)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile("temp/test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	_, err = io.WriteString(f, "Test string")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

}
