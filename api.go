package main

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
	YearStarted  int      `json:"year_started"`
	FirstAlbum   string   `json:"first_album"`
	Members      []string `json:"members"`
}
	
type Location struct {
	// Définissez les champs appropriés
}
	
type Date struct {
	// Définissez les champs appropriés
}

type Relation struct {
	// Définissez les champs appropriés
	}

// Exemple de données (remplacez-les par vos données réelles)
var artistsData = []Artist{
		{Name: "Artist1", Image: "image1.jpg", YearStarted: 2000, FirstAlbum: "Album1", Members: []string{"Member1", "Member2"}},
		{Name: "Artist2", Image: "image2.jpg", YearStarted: 1995, FirstAlbum: "Album2", Members: []string{"Member3", "Member4"}},
		// Ajoutez d'autres données d'artistes
}
	
	// Handler pour l'endpoint des artistes
	func artistsHandler(w http.ResponseWriter, r *http.Request) {
		// Convertissez les données d'artistes en JSON
		jsonData, err := json.Marshal(artistsData)
		if err != nil {
			http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
			return
		}
	
		// Définissez l'en-tête Content-Type pour indiquer que la réponse est du JSON
		w.Header().Set("Content-Type", "application/json")
	
		// Écrivez les données JSON dans la réponse
		w.Write(jsonData)
	}
	
	func main() {
		// Définissez les gestionnaires d'endpoints
		http.HandleFunc("/artists", artistsHandler)
		// Définissez d'autres gestionnaires d'endpoints pour les autres parties de l'API
	
		// Démarrez le serveur sur le port 0000 (ou un autre port de votre choix)
		fmt.Println("Serveur écoutant sur le port 0000...")
		http.ListenAndServe(":0000", nil)
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

    err = json.Unmarshal(body, &artists)
    if err != nil {
        http.Error(w, "Erreur lors de l'analyse JSON", http.StatusInternalServerError)
        return
    }

	renderhtml(w)
}

func renderhtml(w http.ResponseWriter) {
	htmlFile, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier HTML", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("index").Parse(string(htmlFile))
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