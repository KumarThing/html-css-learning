package main

import (
	"encoding/json"
	"os"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))
const dataFile = "review.json"


type Reviews struct{
	Title string
	Author string
	Rating int
	Text string
}

var reviews []Reviews

func main() {
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	loadReviewFromJson()

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


		saveReviewToJson()

		
		
		tmpl.Execute(w, reviews)
		

	})

	fmt.Print("Server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}



func saveReviewToJson() error{
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Error creating JSON file", err)
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(reviews)
}

func loadReviewFromJson() error {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("error loading json file",err)
		return nil
	}

	defer file.Close()

	return json.NewDecoder(file).Decode(&reviews)
}
