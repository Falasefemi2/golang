package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJson(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}
