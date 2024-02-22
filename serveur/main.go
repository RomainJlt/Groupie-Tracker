package main

import (
	"groupie/handler"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler.ArtistHandler)
	// ttp.HandleFunc("/info_artiste", handler.InfoArtiste)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("../styles"))))

	fmt.Println("Server is running")
	
	http.ListenAndServe(":8080", nil)
}
