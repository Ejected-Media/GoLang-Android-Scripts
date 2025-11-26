Yes, absolutely. Go (Golang) works very well in Termux. In fact, because Go compiles to a single binary, it can be extremely fast for file operations on Android.
The setup logic is almost identical to Python, but the code will be a bit more verbose due to Go's explicit error handling.
Here is how to do the exact same file cleanup project using Go in Termux.
1. Install Go in Termux
Open Termux and run:
pkg install golang

2. The Go Script
Create a file named organize.go (you can use nano organize.go inside Termux).
Here is the Go translation of the Python script. It looks for .jpg and .png files in your Downloads folder and moves them to a subfolder.

```
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
```

3. How to Run It
You have two options with Go:
Option A: Run it like a script (Slower, good for testing)
go run organize.go

Option B: Compile it (The "Go way")
This creates a standalone executable file. You could theoretically move this binary to another phone with Termux (and the same architecture) and run it without installing Go again.
go build organize.go
./organize

Important Difference: os.Rename vs shutil.move
In the script above, I used os.Rename.
 * The Catch: os.Rename only works if the source and destination are on the same filesystem (partition).
 * On Android: Since we are moving files from Internal Storage -> Internal Storage, this works perfectly and is instant.
 * If you cross filesystems: If you try to move a file from Internal Storage to an SD Card, os.Rename will fail. You would need to write code to "Copy file -> Delete original."
Would you like me to show you how to modify the script to handle SD cards (Copy+Delete logic), or are we sticking to internal storage for tonight?
