package file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Overwatch01/TermToDo/model"
)

func ReadFile() ([]model.Task, error) {
	filename := getFileName()

	entries, err := readOrCreateJSONFile(filename)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return entries, nil
}

func getFileName() string {
	time := time.Now()
	filename := time.Format("2006-01-02")

	return fmt.Sprintf("./tasks_%v.json", filename)
}

func readOrCreateJSONFile(filename string) ([]model.Task, error) {
	// Check if the file exists
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		// If the file doesn't exist, create it with initial content
		fmt.Println("File does not exist, creating file... " + filename)
		initialEntries := []model.Task{}
		if err := SaveFile(initialEntries); err != nil {
			return nil, fmt.Errorf("error creating JSON file: %w", err)
		}
		return initialEntries, nil
	} else if err != nil {
		// Handle other errors
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// If the file exists, read and parse its content
	var entries []model.Task
	if err := json.NewDecoder(file).Decode(&entries); err != nil {
		return nil, fmt.Errorf("error decoding JSON file: %w", err)
	}

	return entries, nil
}

func SaveFile(entries []model.Task) error {
	// Open the file for writing
	filename := getFileName()
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Serialize the entries to JSON and write to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON
	if err := encoder.Encode(entries); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}
