package main

import (
	"fmt"
	"os"
	"strings"
)

// AsciiArt génère de l'art ASCII pour le texte donné en utilisant le style de bannière spécifié.
func AsciiArt(text, banner string) (string, error) {
	// Par défaut "standard" si la bannière est vide
	if banner == "" {
		banner = "standard"
	}

	// Lire le fichier de bannière
	fontPath := "banners/" + banner + ".txt"
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		return "", fmt.Errorf("failed to read banner file: %w", err)
	}

	// Analyser les données de police
	// Le format standard a généralement 8 lignes par caractère, commençant par l'espace (32) jusqu'à ~ (126).
	// Certains fichiers peuvent avoir un en-tête ou un nombre de lignes différent, mais standard est généralement 8 lignes.
	// Nous supposerons le format standard où chaque caractère a 8 lignes de hauteur.
	// Le fichier commence généralement par une nouvelle ligne, nous devrons peut-être gérer cela.

	lines := strings.Split(string(fontData), "\n")

	// Validation de base du fichier de police
	if len(lines) < 8 {
		return "", fmt.Errorf("invalid banner file format")
	}

	// Mapper les caractères à leur représentation ASCII
	asciiMap := make(map[rune][]string)

	// Les caractères ASCII standard vont de 32 à 126
	// Chaque caractère est de 9 lignes dans le fichier (incluant le séparateur/nouvelle ligne généralement),
	// mais typiquement nous extrayons 8 lignes de contenu.
	// Supposons le format standard :
	// Ligne 0 : vide (parfois)
	// Ligne 1-8 : Espace
	// Ligne 9 : vide
	// Ligne 10-17 : !
	// ...

	// Ajustement pour le format commun :
	// Le fichier contient généralement 95 caractères imprimables (32-126).
	// Chaque caractère a 9 lignes de hauteur ? Ou 8 ?
	// Standard `standard.txt` a généralement 9 lignes par caractère si on compte le séparateur,
	// ou juste 8 lignes d'art.
	// Essayons d'être robustes.
	// Si nous supposons 9 lignes par caractère (1 ligne séparateur vide + 8 lignes art),
	// commençant à la ligne 1 (indexé à 0).

	// Gérons d'abord la normalisation des nouvelles lignes
	var cleanLines []string
	for _, line := range lines {
		// Nous gardons toutes les lignes pour l'instant, mais nous devrons peut-être gérer \r
		cleanLines = append(cleanLines, strings.ReplaceAll(line, "\r", ""))
	}
	lines = cleanLines

	// Si le fichier commence par une ligne vide, la sauter ?
	// En fait, indexons simplement basé sur le code de caractère.
	// Char 32 (espace) commence à la ligne 1 (si 0 est vide) ou 0.
	// Généralement :
	// Ligne 0 : (vide)
	// Ligne 1-8 : Art pour espace
	// Ligne 9 : (vide)
	// Ligne 10-17 : Art pour !

	// Donc char `i` (code ascii) commence à : (i - 32) * 9 + 1

	for i := 32; i <= 126; i++ {
		start := (i-32)*9 + 1
		if start+8 > len(lines) {
			break
		}
		// Extraire 8 lignes
		var charLines []string
		for j := 0; j < 8; j++ {
			charLines = append(charLines, lines[start+j])
		}
		asciiMap[rune(i)] = charLines
	}

	// Générer le résultat
	var result strings.Builder

	// Gérer les nouvelles lignes dans le texte d'entrée
	// Le texte d'entrée peut contenir \r\n ou \n.
	// Nous devrions diviser par \n.
	text = strings.ReplaceAll(text, "\r\n", "\n")
	inputLines := strings.Split(text, "\n")

	for idx, line := range inputLines {
		if line == "" {
			if idx < len(inputLines)-1 {
				result.WriteString("\n")
			}
			continue
		}

		// Pour chaque ligne de texte, nous devons imprimer 8 lignes d'art ASCII
		for row := 0; row < 8; row++ {
			for _, char := range line {
				if art, ok := asciiMap[char]; ok {
					result.WriteString(art[row])
				} else {
					// Gérer les caractères inconnus (optionnel : ne rien imprimer ou un espace réservé)
					// Pour l'instant, sautons ou imprimons un point d'interrogation si nous en avions un,
					// mais nous allons juste ignorer pour éviter de planter.
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
