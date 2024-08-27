package commands

import (
	"fmt"
	"strings"
	"time"
)

func DisplayTime(args []string) {
	if len(args) > 0 && strings.ToLower(args[0]) == "utc" {
		displayUTCTime()
	} else {
		displayLocalTime()
	}
}

func displayLocalTime() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format(time.RFC1123))
}

func displayUTCTime() {
	currentTime := time.Now().UTC()
	fmt.Println(currentTime.Format(time.RFC1123))
}
