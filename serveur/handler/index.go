package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var artists []Artist

type Artist struct {
	Image string `json:"image"`
	Name string `json:"name"`
}

func Index(w http.ResponseWriter, r *http.Request) {
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

	renderIndexHtml(w)
}

func renderIndexHtml(w http.ResponseWriter) {
	htmlFile, err := ioutil.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier HTML", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("../index.html").Parse(string(htmlFile))
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