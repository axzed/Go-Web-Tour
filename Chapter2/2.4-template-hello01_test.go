package Chapter2

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func helloHandleFunc02(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	user := &User{
		Name:   "命运引力",
		Gender: "男",
		Age:    18,
	}
	tmpl.Execute(w, user)
}

func Test06(t *testing.T) {
	http.HandleFunc("/", helloHandleFunc02)
	http.ListenAndServe(":8080", nil)
}
