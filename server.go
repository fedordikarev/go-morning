package main

import (
		"fmt"
		"log"
		"net/http"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/hello", HelloHandler)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/hello", http.StatusFound)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
