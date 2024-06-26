package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = "localhost:7070"

type RedirectServer struct {
	redirectCount int
}

func (s *RedirectServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.redirectCount++
	fmt.Println("Received header: " + req.Header.Get("Known-redirects"))
	http.Redirect(rw, req, fmt.Sprintf("/redirect%d", s.redirectCount), http.StatusTemporaryRedirect)
}

func main() {
	s := http.Server{
		Addr:    addr,
		Handler: &RedirectServer{0},
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := http.Client{}
	redirectCount := 0

	// If the count of redirects is reached
	// than return error.
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		if redirectCount > 2 {
			return fmt.Errorf("too many redirects")
		}
		req.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	_, err := client.Get("http://" + addr)
	if err != nil {
		panic(err)
	}
}
