package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to SnippetBox !!!")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Server listening on port 4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}
