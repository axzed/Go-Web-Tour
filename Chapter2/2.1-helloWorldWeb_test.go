package Chapter2

import (
	"fmt"
	"net/http"
	"testing"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func Test1(t *testing.T) {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
