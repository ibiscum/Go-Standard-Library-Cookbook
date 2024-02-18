package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received form data: %v\n", req.Form)
	fmt.Printf("Received header: %v\n", req.Header)

	_, err = rw.Write([]byte(string(s)))
	if err != nil {
		log.Fatal(err)
	}
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	form := url.Values{}
	form.Set("id", "5")
	form.Set("name", "Wolfgang")

	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:7070",
		strings.NewReader(form.Encode()))

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type",
		"application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	res.Body.Close()
	fmt.Println("Response from server:" + string(data))

}
