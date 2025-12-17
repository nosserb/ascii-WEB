package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var templates *template.Template

func init() {
	// Analyser tous les templates
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	// Ajouter aussi index.html de la racine
	if _, err := templates.ParseGlob("index.html"); err != nil {
		log.Fatalf("Error parsing index.html: %v", err)
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

	r.ParseForm()
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Nettoyer les \r
	text = strings.ReplaceAll(text, "\r", "")

	if text == "" {
		errorHandler(w, http.StatusBadRequest)
		return
	}

	result, err := AsciiArt(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		errorHandler(w, http.StatusInternalServerError)
		return
	}

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
