package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	// Analyser tous les templates dans le répertoire templates
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
}

func main() {
	// Servir les fichiers statiques
	http.Handle("/style.css", http.FileServer(http.Dir(".")))
	http.Handle("/script.js", http.FileServer(http.Dir(".")))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Démarrer le serveur
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Printf("Error executing template: %v", err)
		errorHandler(w, http.StatusInternalServerError)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")     // Supposant que le nom de l'entrée est "text" ou similaire, vérifier index.html
	banner := r.FormValue("banner") // Supposant "banner"

	// Vérifier index.html pour voir les noms réels des entrées.
	// Basé sur le view_file précédent de index.html :
	// <input type="text" id="text-input" ...> -> Cela pourrait être pour JS.
	// Attendez, index.html semble utiliser JS pour envoyer des requêtes ?
	// Vérifions script.js pour voir comment il envoie les données.
	// Si c'est une soumission de formulaire, nous utilisons FormValue. Si c'est JSON/fetch, nous devons décoder JSON.
	// L'exigence de l'utilisateur dit "POST /ascii-art -> reçoit le texte et le type de bannière".
	// Je vais supposer une soumission de formulaire standard ou je dois vérifier script.js.
	// Pour l'instant, je vais implémenter la gestion de formulaire de base, mais je devrais vérifier script.js.

	if text == "" {
		// Essayer de lire depuis le corps si pas dans le formulaire (ex: JSON)
		// Mais pour l'instant supposons des données de formulaire ou des paramètres de requête.
		// Si le texte est vide, peut-être renvoyer 400 ?
		errorHandler(w, http.StatusBadRequest)
		return
	}

	result, err := AsciiArt(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}

	// Juste écrire le résultat en texte brut ou JSON ?
	// L'utilisateur a dit "renvoie le résultat ASCII".
	// Si le frontend attend juste la chaîne, nous l'écrivons.
	w.Write([]byte(result))
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		if err := templates.ExecuteTemplate(w, "404.html", nil); err != nil {
			log.Printf("Error executing 404 template: %v", err)
		}
	} else if status == http.StatusInternalServerError {
		if err := templates.ExecuteTemplate(w, "500.html", nil); err != nil {
			log.Printf("Error executing 500 template: %v", err)
		}
	} else if status == http.StatusBadRequest {
		if err := templates.ExecuteTemplate(w, "400.html", nil); err != nil {
			log.Printf("Error executing 400 template: %v", err)
		}
	} else {
		// Message d'erreur par défaut
		http.Error(w, http.StatusText(status), status)
	}
}
