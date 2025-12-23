package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		tmpl.Execute(w, nil)
	})

	fmt.Print("Server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}