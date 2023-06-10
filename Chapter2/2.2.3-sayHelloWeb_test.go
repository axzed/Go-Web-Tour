package Chapter2

import (
	"net/http"
	"testing"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func Test2(t *testing.T) {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", nil)
}
