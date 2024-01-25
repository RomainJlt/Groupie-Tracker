package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)
var artists []Artist

type Artist struct {
	Image string `json:"image"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":8080", nil)
}

func server(w http.ResponseWriter, r *http.Request) {
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

	var artists []Artist
    err = json.Unmarshal(body, &artists)
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
        return
    }

	for _, artist := range artists {
		fmt.Println(artist.Name)
	}	
}
 