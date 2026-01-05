package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getFloatFromFile(fileName string) (float64, error) {
	// Placeholder function to simulate reading balance from a file
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file")
	}
	
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored e")
	}
	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	// Placeholder function to simulate writing balance to a file
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}