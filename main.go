package main

import (
	"fmt"
	"os"

	"github.com/shaheemMPM/mac-tricks/commands"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "cc-stremio" {
		commands.ClearStremioCache()
		return
	}

	fmt.Println("Available options:")
	fmt.Println("1. Clear Stremio cache (cc-stremio)")

	var option string
	fmt.Print("Enter option number: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		commands.ClearStremioCache()
	default:
		fmt.Println("Invalid option")
	}
}
