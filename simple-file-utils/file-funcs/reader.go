package fileUtils

import (
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	// Fetch the user's home directory
	homeDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get user home directory: %v", err)
	}

	// Construct the full path to the file
	fullPathToFile := fmt.Sprintf("%s/%s", homeDir, filename)

	// Read the file's contents
	contents, err := os.ReadFile(fullPathToFile)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Return the file's contents as a string
	return string(contents), nil
}
