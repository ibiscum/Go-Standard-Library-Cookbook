package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Simply write string
	_, err := io.WriteString(os.Stdout, "This is string to standard output.\n")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(os.Stderr, "This is string to standard error output.\n")
	if err != nil {
		log.Fatal(err)
	}

	// Stdout/err implements
	// writer interface
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 20; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// The fmt package
	// could be used too
	fmt.Fprintln(os.Stdout)
}
