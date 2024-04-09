package main

import (
	"io"
	"log"
	"net/http"

	c "github.com/sjtiffin/gotourney/ctrl"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /brackets", getBracketsHandler)
	mux.HandleFunc("GET /brackets/{id}", getBracketByIDHandler)

	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

func getBracketByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	io.WriteString(w, "you have asked for "+id)
}

func getBracketsHandler(w http.ResponseWriter, r *http.Request) {
	bs := c.Brackets()
	_, err := w.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
}
