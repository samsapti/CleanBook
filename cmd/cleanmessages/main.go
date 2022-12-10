package main

import (
	"flag"
	"os"

	"github.com/samsapti/CleanMessages/internal/conversation"
	"github.com/samsapti/CleanMessages/internal/utils"
)

var (
	// currently just a JSON file path
	path *string = flag.String("path", "", "Path to the directory containing your Facebook data (required)")
	port *int    = flag.Int("port", 8080, "Port to listen on")
)

func main() {
	flag.Parse()

	if len(*path) == 0 {
		utils.PrintError("error: -path must be specified\n")
		flag.Usage()
		os.Exit(1)
	}

	conv, err := conversation.Parse(*path)
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
