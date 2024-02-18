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
	rw.Write([]byte(string(s)))
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

	useRequest()
	simplePost()

}

func simplePost() {
	res, err := http.Post("http://localhost:7070",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Radek&surname=Sohlich"))
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

func useRequest() {

	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "Radek")
	form.Add("surname", "Sohlich")

	req, err := http.NewRequest("POST", "http://localhost:7070", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := hc.Do(req)
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
