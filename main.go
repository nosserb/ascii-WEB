package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	// Middleware pour les headers CORS
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	})

	// Serveur de fichiers statiques pour tous les fichiers
	mux.Handle("/", http.FileServer(http.Dir(".")))

	// Route pour l'API ASCII art
	mux.HandleFunc("/ascii-art", asciiArtHandler)

	// Démarrer le serveur
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Lire le body brut d'abord
	bodyBytes, _ := io.ReadAll(r.Body)
	log.Printf("Raw body received: %q", string(bodyBytes))

	// Décoder le JSON
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	var payload struct {
		Text   string `json:"text"`
		Banner string `json:"banner"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Bad request - invalid JSON", http.StatusBadRequest)
		return
	}

	text := payload.Text
	banner := payload.Banner

	log.Printf("Decoded - text: %q, banner: %q", text, banner)

	// Nettoyer les \r
	text = strings.ReplaceAll(text, "\r", "")

	if text == "" {
		log.Printf("Text is empty")
		http.Error(w, "Bad request - empty text", http.StatusBadRequest)
		return
	}

	result, err := AsciiArt(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
