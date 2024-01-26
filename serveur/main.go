package main

import (
	"groupie/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.Handle("./styles/", http.StripPrefix("./styles/", http.FileServer(http.Dir("styles.css"))))
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/info_artiste", handler.InfoArtiste)
	http.Handle("./styles/", http.StripPrefix("./styles/", http.FileServer(http.Dir("info_atiste.css"))))

	http.HandleFunc("/index", handler.InfoArtiste)
	http.Handle("./styles/", http.StripPrefix("./styles/", http.FileServer(http.Dir("info_atiste.css"))))

}
