package banner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ASCII art banner character constants
const (
	// ASCIIStart is the first ASCII character code (space)
	ASCIIStart = 32
	// ASCIIEnd is the last ASCII character code (tilde)
	ASCIIEnd = 126
	// TotalChars is the total number of supported ASCII characters
	TotalChars = 95
	// LinesPerChar is the height in lines of each ASCII art character
	LinesPerChar = 8
)

// LoadBanner reads an ASCII art banner file and parses it into a map.
// Each character (from ASCII 32 to 126) is represented by 8 lines of text.
// The banner file should contain 9 lines per character (8 lines + 1 separator).
// Returns a map where each rune is associated with its 8-line representation.
func LoadBanner(path string) (map[rune][]string, error) {
	// Read all lines from the banner file
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}


	// Add an empty line at the end if it's missing (some files may not have the final newline)
	expectedLines := TotalChars * 9
	if len(lines) == expectedLines-1 {
		lines = append(lines, "")
	}

	// Validate the total number of lines in the file
	if len(lines) < expectedLines-1 {
		return nil, fmt.Errorf("invalid line count: expected at least %d, got %d", expectedLines-1, len(lines))
	}

	// Create a map to store the ASCII art for each character
	banner := make(map[rune][]string)

	// Parse each character's representation from the file
	for i := 0; i < TotalChars; i++ {
		// Calculate which character this is (A=65, etc.)
		char := rune(ASCIIStart + i)
		// Calculate where this character's lines start in the file
		startIndex := i * 9

		// Extract the 8 lines for this character (skip the first empty line)
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
// It properly handles both Unix (\n) and Windows (\r\n) line endings.
func readLines(path string) ([]string, error) {
	// Open the banner file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		// Remove carriage return characters (Windows line ending)
		line = strings.TrimRight(line, "\r")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
