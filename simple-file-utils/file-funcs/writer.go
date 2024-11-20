package fileUtils

import (
	"os"
	"fmt"
)

func WriteToFile(filename string, data string) error{
	// Fetch the user's home directory
	homeDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get user home directory: %v", err)
	}
	
	// Construct the full path to the file
	fullPathToFile := fmt.Sprintf("%s/%s", homeDir, filename)

	// Convert incoming data to a byte array and write to file with correct permissions
	err = os.WriteFile(fullPathToFile, []byte(data), 0666)
	if err != nil {
		return err
	}

	return nil
}