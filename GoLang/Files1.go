package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 1. Define paths
	// In Termux, the home directory is usually /data/data/com.termux/files/home
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error finding home directory:", err)
		return
	}

	// This path relies on you having run 'termux-setup-storage' previously
	sourceDir := filepath.Join(homeDir, "storage", "downloads")
	destDir := filepath.Join(sourceDir, "Sorted_Images_Go")

	// 2. Create the destination folder if it doesn't exist
	// 0755 is the standard permission (read/write/execute for owner, read/execute for others)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// 3. Read the source directory
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Println("Error reading source directory:", err)
		return
	}

	// 4. Loop through files and move them
	for _, file := range files {
		// Skip if it's a directory
		if file.IsDir() {
			continue
		}

		name := file.Name()
		lowerName := strings.ToLower(name)

		// Check for image extensions
		if strings.HasSuffix(lowerName, ".jpg") || strings.HasSuffix(lowerName, ".png") {
			oldPath := filepath.Join(sourceDir, name)
			newPath := filepath.Join(destDir, name)

			// Move the file
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("Failed to move %s: %v\n", name, err)
			} else {
				fmt.Printf("Moved: %s\n", name)
			}
		}
	}
}
