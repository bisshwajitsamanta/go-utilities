package main

import (
	"fmt"
	"net/http"
)

const Port = ":8080"

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Starting Application on %s", Port))
	_ = http.ListenAndServe(Port, nil)
}
