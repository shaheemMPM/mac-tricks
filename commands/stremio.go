package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/shaheemMPM/mac-tricks/helpers"
)

const StremioCachePath = "~/Library/Application Support/stremio-server/stremio-cache"

func ClearStremioCache() {
	expandedPath, err := helpers.ExpandTilde(StremioCachePath)
	if err != nil {
		fmt.Println("Error expanding path:", err)
		return
	}

	if _, err := os.Stat(expandedPath); os.IsNotExist(err) {
		fmt.Println("Stremio cache directory does not exist")
		return
	}

	size, err := helpers.GetDirSize(expandedPath)
	if err != nil {
		fmt.Println("Error calculating directory size:", err)
		return
	}

	sizeInMB := float64(size) / (1024 * 1024)
	if sizeInMB >= 1024 {
		sizeInGB := sizeInMB / 1024
		fmt.Printf("Total size of Stremio cache: %.2f GB\n", sizeInGB)
	} else {
		fmt.Printf("Total size of Stremio cache: %.2f MB\n", sizeInMB)
	}

	var confirm string
	fmt.Print("Do you want to delete this cache? (y/n): ")
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) == "y" {
		err = os.RemoveAll(expandedPath)
		if err != nil {
			fmt.Println("Error deleting cache:", err)
			return
		}
		fmt.Println("Stremio cache cleared successfully")
	} else {
		fmt.Println("Operation cancelled")
	}
}
