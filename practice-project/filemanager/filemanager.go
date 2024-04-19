package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Open the file named "prices.txt"
	file, err := os.Open(path)

	// Check if there's an error opening the file
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	var lines []string

	// Read each line from the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check if there's an error during scanning
	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}
