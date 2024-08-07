package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello My nameis Nishat Wasi and I am eager to learn GO lang!!! ")

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parseforn error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("fname")
	lname := r.FormValue("lname")
	fmt.Fprintf(w, "fname =%s\n", name)
	fmt.Fprintf(w, "lname=%s\n", lname)

}

func main() {

	FileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", FileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("starting server at port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
