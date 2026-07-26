package fileops

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func WriteFloatToFile(balance float64, fileName string) error {
	balanceText := fmt.Sprint(balance)
	filePath := filepath.Join(".", fileName)
	err := os.WriteFile(filePath, []byte(balanceText), 0644)
	if err != nil {
		return fmt.Errorf("failed to write balance to file: %w", err)
	}

	return nil
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to read a float from file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse file value")
	}
	return value, nil
}
