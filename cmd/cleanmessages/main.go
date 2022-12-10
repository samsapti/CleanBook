package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/samsapti/CleanMessages/internal/utils"
	"github.com/samsapti/CleanMessages/pkg/conversation"
)

var (
	basePath *string = flag.String("path", "", "Path to the directory containing your Facebook data (required)")
	port     *int    = flag.Int("port", 8080, "Port to listen on")

	conversations []*conversation.Conversation
)

func main() {
	flag.Parse()

	// Quit if -path was not specified
	if len(*basePath) == 0 {
		flag.Usage()
		utils.PrintPanic("error: -path must be specified\n")
	}

	inbox := filepath.Join(*basePath, "messages", "inbox")
	convDirs, err := os.ReadDir(inbox)
	if err != nil {
		utils.PrintError("error: failed to open messages directory: %s", err)
		os.Exit(1)
	}

	for _, v := range convDirs {
		if !v.IsDir() {
			continue
		}

		filePath := filepath.Join(inbox, v.Name(), "message_1.json")
		conv, err := conversation.Parse(filePath)
		if err != nil {
			utils.PrintError("error: could not parse JSON file: %s", err)
		}

		conversations = append(conversations, conv)
	}
}
