package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /brackets/{id}", getBracketByIDHandler)

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

func getBracketByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	io.WriteString(w, "you have asked for "+id)
}
