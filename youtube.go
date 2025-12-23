package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("youtube_clone/index.html"))

func main() {
	http.Handle("/youtube_clone/static/", http.StripPrefix("/youtube_clone/static/", http.FileServer(http.Dir("youtube_clone/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tmpl.Execute(w, nil)
	})

	fmt.Println("your server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)


}