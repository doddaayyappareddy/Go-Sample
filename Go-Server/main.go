package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "Status Not Found", err)
		return
	}
	fmt.Fprintf(res, "Post request succes")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Printf(name, address)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "Status Not Found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Method Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("startung server at port no 9080 \n ")
	if err := http.ListenAndServe(":9080", nil); err != nil {
		log.Fatal(err)
	}
}
