package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type Artist struct {
	Image string `json:"image"`
	Name string `json:"name"`
}

func renderIndexHtml(w http.ResponseWriter, data []Artist) {
	tmpl, err := template.ParseFiles("../index.html")
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du modèle HTML", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML: %v", err), http.StatusInternalServerError)
		return
	}
}
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    url := "https://groupietrackers.herokuapp.com/api/artists"
    resp, err := http.Get(url)
	
    if err != nil {
        http.Error(w, "Erreur lors de la requête HTTP", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusInternalServerError)
        return
    }

    var artists []Artist
    err = json.Unmarshal(body, &artists)
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
        return
    }

    var filteredArtists []Artist

    for _, artist := range artists {
        if query != "" {
            if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
                filteredArtists = append(filteredArtists, artist)
            }
        } else {
            filteredArtists = append(filteredArtists, artist)
        }
    }

	var artistsData []Artist
    for _, artist := range filteredArtists {
        data := Artist{
            Image:        artist.Image,
            Name:         artist.Name,
        }
        artistsData = append(artistsData, data)
    }

	renderIndexHtml(w, artistsData)
}