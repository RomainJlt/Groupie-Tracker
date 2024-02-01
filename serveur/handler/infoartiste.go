package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
		return
	}

	renderIndexHtml(w)
}
func Locations(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
		return
	}

	renderIndexHtml(w)
}
func Relation(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
		return
	}

	renderIndexHtml(w)
}
func Dates(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
		return
	}

	renderIndexHtml(w)
}
