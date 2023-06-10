package Chapter2

import (
	"net/http"
	"testing"
)

type Refer struct {
	handler http.Handler
	refer   string
}

func (r *Refer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Referer() == r.refer {
		r.handler.ServeHTTP(resp, req)
	} else {
		resp.WriteHeader(http.StatusForbidden)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func Test03(t *testing.T) {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.axzed.com",
	}
	http.HandleFunc("/hello2", hello2)
	http.ListenAndServe(":8080", referer)
}
