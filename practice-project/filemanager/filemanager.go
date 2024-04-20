package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	// Open the file named "prices.txt"
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
