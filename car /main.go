package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var homeTml = template.Must(template.ParseFiles("template/home.html"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/api/img/", http.StripPrefix("/api/img/", http.FileServer(http.Dir("api/img"))))

	http.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request){
		homeTml.Execute(w, nil)
	})

	fmt.Println("Your server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}