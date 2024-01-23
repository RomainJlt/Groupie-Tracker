package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Artist struct {
	Image string `json:"image"`
	Name string `json:"name"`
}

func main() {
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

	fmt.Println(artists)
}