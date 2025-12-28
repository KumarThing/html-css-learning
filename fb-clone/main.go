package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("fb-clone.html"))

func main() {
	
	http.Handle("/fb-clone/static/", http.StripPrefix("/fb-clone/static/", http.FileServer(http.Dir("fb-clone/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tmpl.Execute(w, nil)
	})

	fmt.Println("Server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	
}