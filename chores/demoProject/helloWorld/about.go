package main

import (
	"fmt"
	"net/http"
)

const PortNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is Home Page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of Bytes Written to console"), n)
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := calculateValue(2, 2)
	n, err := fmt.Fprintf(w, "This is About Page and 2 +2 is %d", sum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of Bytes Written to console"), n)
}
func calculateValue(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Application on %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
