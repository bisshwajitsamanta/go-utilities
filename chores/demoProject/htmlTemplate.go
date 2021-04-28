package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const Port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
	//_,_ = fmt.Fprintf(w,"Hello World!")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting Application on %s", Port))
	_ = http.ListenAndServe(Port, nil)
}
