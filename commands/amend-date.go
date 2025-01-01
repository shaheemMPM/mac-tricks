package commands

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func AmendDate(args []string) {
	var targetDate time.Time
	var err error

	// If no date is provided, use current time
	if len(args) == 0 {
		targetDate = time.Now()
	} else {
		// Parse the provided date string
		// Expected format: "2024-12-24 08:06:51"
		targetDate, err = time.Parse("2006-01-02 15:04:05", args[0])
		if err != nil {
			fmt.Println("Error: Invalid date format. Please use: YYYY-MM-DD HH:MM:SS")
			fmt.Println("Example: amend-date \"2024-12-24 08:06:51\"")
			return
		}
	}

	// Format the date string
	dateStr := targetDate.Format("2006-01-02 15:04:05")

	// Set up the git command
	cmd := exec.Command("git", "commit", "--amend", "--date="+dateStr, "--no-edit")

	// Set the GIT_COMMITTER_DATE environment variable
	cmd.Env = append(os.Environ(), "GIT_COMMITTER_DATE="+dateStr)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error amending commit: %v\n%s", err, output)
		return
	}

	fmt.Printf("Successfully amended commit date to: %s\n", dateStr)
}
