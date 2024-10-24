package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shaheemMPM/mac-tricks/helpers"
)

type NodeModulesInfo struct {
	Paths     []string
	TotalSize int64
}

func CleanNodeModules() {
	info, err := findNodeModules(".")
	if err != nil {
		fmt.Println("Error scanning directories:", err)
		return
	}

	if len(info.Paths) == 0 {
		fmt.Println("No node_modules directories found")
		return
	}

	fmt.Printf("Found %d node_modules directories:\n", len(info.Paths))
	for _, path := range info.Paths {
		size, err := helpers.GetDirSize(path)
		if err != nil {
			fmt.Printf("- %s (error calculating size)\n", path)
		} else {
			sizeInMB := float64(size) / (1024 * 1024)
			if sizeInMB >= 1024 {
				fmt.Printf("- %s (%.2f GB)\n", path, sizeInMB/1024)
			} else {
				fmt.Printf("- %s (%.2f MB)\n", path, sizeInMB)
			}
		}
	}

	totalSizeInMB := float64(info.TotalSize) / (1024 * 1024)
	if totalSizeInMB >= 1024 {
		fmt.Printf("\nTotal size: %.2f GB\n", totalSizeInMB/1024)
	} else {
		fmt.Printf("\nTotal size: %.2f MB\n", totalSizeInMB)
	}

	var confirm string
	fmt.Print("\nDo you want to delete all these node_modules directories? (y/n): ")
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) == "y" {
		for _, path := range info.Paths {
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Printf("Error deleting %s: %v\n", path, err)
			} else {
				fmt.Printf("Deleted: %s\n", path)
			}
		}
		fmt.Println("\nCleanup completed!")
	} else {
		fmt.Println("Operation cancelled")
	}
}

func findNodeModules(root string) (*NodeModulesInfo, error) {
	result := &NodeModulesInfo{
		Paths:     make([]string, 0),
		TotalSize: 0,
	}

	err := filepath.Walk(root, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			return nil
		}

		if fileInfo.Name() == "node_modules" {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			size, err := helpers.GetDirSize(absPath)
			if err != nil {
				return err
			}

			result.Paths = append(result.Paths, absPath)
			result.TotalSize += size

			return filepath.SkipDir
		}

		return nil
	})

	return result, err
}
