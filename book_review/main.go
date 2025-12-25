package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))


type Reviews struct{
	Title string
	Author string
	Rating int
	Text string
}

var reviews []Reviews

func main() {
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		tmpl.Execute(w, reviews)
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		r.ParseForm()

		rating, _ := strconv.Atoi(r.FormValue("rating"))
		review := Reviews {
			Title: r.FormValue("title"),
			Author : r.FormValue("author"),
			Rating: rating,
			Text : r.FormValue("text"),

		}
		reviews = append(reviews, review)

		
		
		tmpl.Execute(w, reviews)
		

	})

	fmt.Print("Server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}