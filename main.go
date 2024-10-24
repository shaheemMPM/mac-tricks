package main

import (
	"fmt"
	"os"

	"github.com/shaheemMPM/mac-tricks/commands"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "cc-stremio":
			commands.ClearStremioCache()
		case "git-tree":
			commands.DisplayGitAwareTree()
		case "time":
			commands.DisplayTime(os.Args[2:])
		case "clean-nm":
			commands.CleanNodeModules()
		default:
			fmt.Println("Unknown command")
		}
		return
	}

	fmt.Println("Available options:")
	fmt.Println("1. Clear Stremio cache (cc-stremio)")
	fmt.Println("2. Show git file tree (git-tree)")
	fmt.Println("3. Display current time (time)")
	fmt.Println("4. Clean node_modules directories (clean-nm)")

	var option string
	fmt.Print("Enter option number: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		commands.ClearStremioCache()
	case "2":
		commands.DisplayGitAwareTree()
	case "3":
		commands.DisplayTime(nil)
	case "4":
		commands.CleanNodeModules()
	default:
		fmt.Println("Invalid option")
	}
}
