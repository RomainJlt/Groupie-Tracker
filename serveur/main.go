package main

import (
	"groupie/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles")))) // remplace les ------------- par le chemin du dossier contenant les fichiers css
	http.ListenAndServe(":8080", nil)
}