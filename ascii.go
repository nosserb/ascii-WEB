package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ASCII art banner character constants
const (
	ASCIIStart   = 32
	ASCIIEnd     = 126
	TotalChars   = 95
	LinesPerChar = 8
)

// loadBanner reads an ASCII art banner file and parses it into a map.
func loadBanner(path string) (map[rune][]string, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	expectedLines := TotalChars * 9
	if len(lines) == expectedLines-1 {
		lines = append(lines, "")
	}

	if len(lines) < expectedLines-1 {
		return nil, fmt.Errorf("invalid line count: expected at least %d, got %d", expectedLines-1, len(lines))
	}

	banner := make(map[rune][]string)

	for i := 0; i < TotalChars; i++ {
		char := rune(ASCIIStart + i)
		startIndex := i * 9

		charLines := make([]string, LinesPerChar)
		for j := 0; j < LinesPerChar; j++ {
			if startIndex+j+1 < len(lines) {
				charLines[j] = lines[startIndex+j+1]
			}
		}

		banner[char] = charLines
	}

	return banner, nil
}

// readLines reads all lines from a file and returns them as a slice of strings.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\r")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// printASCII converts input text into ASCII art using the provided font map.
func printASCII(text string, font map[rune][]string) (string, error) {
	const lineHeight = 8

	if text == "" {
		return "", nil
	}

	lines := strings.Split(text, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for asciiLineIdx := 0; asciiLineIdx < lineHeight; asciiLineIdx++ {
			for _, char := range line {
				charLines, exists := font[char]
				if !exists {
					// Remplacer les caractères non supportés par un espace
					char = ' '
					charLines, _ = font[char]
				}

				if len(charLines) < lineHeight {
					return "", fmt.Errorf("invalid font data for character %q: expected %d lines, got %d", char, lineHeight, len(charLines))
				}

				result.WriteString(charLines[asciiLineIdx])
			}
			result.WriteString("\n")
		}
	}

	output := strings.TrimSuffix(result.String(), "\n")
	return output, nil
}

// AsciiArt génère l'art ASCII à partir du texte et du style de bannière fournis.
// Si bannerName est vide, utilise "standard" par défaut.
// Retourne le résultat ASCII art ou une erreur.
func AsciiArt(text, bannerName string) (string, error) {
	// Replace literal \n with actual newlines
	text = strings.ReplaceAll(text, "\\n", "\n")

	// Default banner style
	if bannerName == "" {
		bannerName = "standard"
	}

	// Construire le chemin vers le fichier de bannière
	bannerPath := "banners/" + bannerName

	// Load the banner font file
	asciiFont, err := loadBanner(bannerPath)
	if err != nil {
		return "", fmt.Errorf("error loading banner: %v", err)
	}

	// Render the input text as ASCII art
	result, err := printASCII(text, asciiFont)
	if err != nil {
		return "", fmt.Errorf("error during rendering: %v", err)
	}

	return result, nil
}
