package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
