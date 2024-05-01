package main

import (
	"io"
	"log"
	"net/http"
)

type User string

func (u User) toString() string {
	return string(u)
}

type AuthHandler func(u User, w http.ResponseWriter, r *http.Request)

func main() {

	// Secured API
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Secure(func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w,
			`[{"id":"1","login":"ffghi"},{"id":"2","login":"ffghj"}]`)
		if err != nil {
			log.Fatal(err)
		}
	}))

	mux.HandleFunc("/api/profile", WithUser(func(u User, w http.ResponseWriter, r *http.Request) {
		log.Println(u.toString())
		_, err := io.WriteString(w, "{\"user\":\""+u.toString()+"\"}")
		if err != nil {
			log.Fatal(err)
		}
	}))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func WithUser(h AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-User")
		if len(user) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(User(user), w, r)
	}

}

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sec := r.Header.Get("X-Auth")
		if sec != "authenticated" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(w, r) // use the handler
	}

}
