package main

import (
	"fmt"
	"log"
)

func main() {

	var name string
	fmt.Println("What is your name?")
	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		log.Fatal(err)
	}

	var age int
	fmt.Println("What is your age?")
	_, err = fmt.Scanf("%d\n", &age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s, your age is %d\n", name, age)

}
