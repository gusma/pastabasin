package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to pastabasin'. Your own personal pastebin."))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

	log.Print("Server is running on port 4000")

	err:= http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}