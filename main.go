package main

import (
	"log"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Nathan! Greetings."))
	if err != nil {
		return
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomeHandler)

	log.Println("Starting server on http://localhost:5040")
	err := http.ListenAndServe(":5040", mux)
	log.Fatal(err)

}
