package quote

import (
	"encoding/json"
	"fmt"
	"os"
)

type Quote struct {
	Q string
	A string
	c string
	h string
}

func GetQuotes() ([]Quote, error) {
	// Check if the file exists
	file, err := os.Open("quotes.json")
	if os.IsNotExist(err) {
		return []Quote{}, nil
	} else if err != nil {
		// Handle other errors
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// If the file exists, read and parse its content
	var quotes []Quote
	if err := json.NewDecoder(file).Decode(&quotes); err != nil {
		return nil, fmt.Errorf("error decoding JSON file: %w", err)
	}

	return quotes, nil
}
