package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	var email string
	fmt.Println("Enter username for smtp: ")
	_, err := fmt.Scanln(&email)
	if err != nil {
		log.Fatal(err)
	}

	var pass string
	fmt.Println("Enter password for smtp: ")
	_, err = fmt.Scanln(&pass)
	if err != nil {
		log.Fatal(err)
	}

	auth := smtp.PlainAuth("",
		email,
		pass,
		"smtp.gmail.com")

	c, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	config := &tls.Config{ServerName: "smtp.gmail.com"}

	if err = c.StartTLS(config); err != nil {
		panic(err)
	}

	if err = c.Auth(auth); err != nil {
		panic(err)
	}

	if err = c.Mail(email); err != nil {
		panic(err)
	}
	if err = c.Rcpt(email); err != nil {
		panic(err)
	}

	w, err := c.Data()
	if err != nil {
		panic(err)
	}

	msg := []byte("Hello this is content")
	if _, err := w.Write(msg); err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		panic(err)
	}
	err = c.Quit()

	if err != nil {
		panic(err)
	}

}
