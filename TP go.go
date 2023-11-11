package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	temp, _ := template.ParseFiles("index.html")

	type Promo struct {
		Nom     string
		Filiere string
		Niveau  int
		Nombre  int
	}

	type Etudiant struct {
		Nom    string
		Prenom string
		Sexe   int
		Age    int
	}
	    

	type PageData struct {
		Promo         Promo
		ListeEtudiant []Etudiant
	}

	ListeEtudiant := []Etudiant{{"RODRIGUES", "Cyril", 1, 22}, {"MEDERREG", "Kheir-eddine", 1, 22}, {"PHILIPIERT", "Alan", 1, 22}}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		dataPage := PageData{
			Promo{"Mentor'ac",
            "Informatique",
            5,
            3}, ListeEtudiant}

		temp.ExecuteTemplate(w, "index.html", dataPage)
	})

	fmt.Println("localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
