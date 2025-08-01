package utils

import (
	"fmt"
	"os"
)

// AppendToFile opens an existing file and appends the given data to it.
// It writes a newline before the new content for clarity.
// If the file doesn't exist, it returns an error.
func AppendToFile(path string, data []byte) error {
	// Open file in append mode (write-only, must exist)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	// Write a newline before appending the actual data
	if _, err := file.Write([]byte("\n")); err != nil {
		return fmt.Errorf("cannot write newline: %w", err)
	}

	// Write the new content to the end of the file
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}

	return nil
}
