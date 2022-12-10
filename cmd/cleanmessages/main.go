package main

import (
	"flag"
	"os"

	"github.com/samsapti/CleanMessages/internal/conversation"
	"github.com/samsapti/CleanMessages/internal/utils"
)

var (
	port    int
	dirPath string // currently just a JSON file path
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&dirPath, "path", "", "Path to the directory containing your Facebook data (required)")
	flag.Parse()

	if len(dirPath) == 0 {
		utils.PrintError("You must provide the path to your messages directory")
		flag.Usage()
		os.Exit(1)
	}

	conv, err := conversation.Parse(dirPath)
	if err != nil {
		panic("Error parsing JSON. Aborting...")
	}

	for i, s := range conv.Participants {
		utils.PrintInfo("Participant %d: %s", i, s.Name)
	}

	for _, m := range conv.Messages {
		utils.PrintInfo("%s: %s", m.SenderName, m.Content)
	}
}
