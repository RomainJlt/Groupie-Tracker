package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var artists []Artist
var filteredartists []Artist

type Artist struct {
	Image string `json:"image"`
	Name string `json:"name"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err); return
	}

    err = json.Unmarshal(body, &artists)
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
        return
    }

	for _, artist := range artists {
		if query != "" {
			if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
				filteredartists = append(filteredartists, artist)
			}
		} else {
			filteredartists = append(filteredartists, artist)
		}
	}

	renderIndexHtml(w)
}

func renderIndexHtml(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("../index.html")
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du modèle HTML", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML: %v", err), http.StatusInternalServerError)
		return
	}
}