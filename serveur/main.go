package main

import (
	"groupie/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/info_artiste", handler.InfoArtiste)
	http.HandleFunc("/index", handler.Index)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("../styles"))))
	
	http.ListenAndServe(":8080", nil)

	
}
