package render

import (
	"fmt"
	"strings"
)

// PrintASCII converts input text into ASCII art using the provided font map.
// It handles multi-line text (\n) and renders each character horizontally.
// Each character is 8 lines tall. Empty lines result in 8 blank lines.
// Returns the rendered ASCII art as a string, or an error if unsupported characters are found.
func PrintASCII(text string, font map[rune][]string) (string, error) {
	// Height of each ASCII art character
	const lineHeight = 8

	// Handle empty input
	if text == "" {
		return "", nil
	}

	// Split input by newlines to handle multi-line text
	lines := strings.Split(text, "\n")

	// Build the output string
	var result strings.Builder

	// Process each line of input text
	for _, line := range lines {
		// Empty line produces 1 blank line
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Render each of the 8 horizontal lines for this text line
		for asciiLineIdx := 0; asciiLineIdx < lineHeight; asciiLineIdx++ {
			// For each character in the input line
			for _, char := range line {
				// Look up the ASCII art representation for this character
				charLines, exists := font[char]
				if !exists {
					return "", fmt.Errorf("unsupported character: %q", char)
				}

				// Validate the font data has the correct number of lines
				if len(charLines) < lineHeight {
					return "", fmt.Errorf("invalid font data for character %q: expected %d lines, got %d", char, lineHeight, len(charLines))
				}

				// Append the current line of this character's art
				result.WriteString(charLines[asciiLineIdx])
			}
			// Add newline after completing a horizontal line
			result.WriteString("\n")
		}
	}

	// Remove the trailing newline from the final output
	output := strings.TrimSuffix(result.String(), "\n")

	return output, nil
}
