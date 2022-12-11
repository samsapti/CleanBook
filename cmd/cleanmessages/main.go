package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/samsapti/CleanMessages/internal/utils"
	"github.com/samsapti/CleanMessages/pkg/conversation"
)

var (
	basePath *string = flag.String("d", "", "Path to the directory containing your Facebook data (required)")
	port     *int    = flag.Int("p", 8080, "Port to listen on")

	convs []*conversation.Conversation
)

func main() {
	flag.Parse()

	// Quit if -path was not specified
	if len(*basePath) == 0 {
		flag.Usage()
		utils.PrintFatal("error: -path must be specified\n")
	}

	messagesPath := filepath.Join(*basePath, "messages")
	inboxPath := filepath.Join(messagesPath, "inbox")
	convDirs, err := os.ReadDir(inboxPath)
	if err != nil {
		utils.PrintFatal("error: failed to open messages directory: %s", err)
	}

	for _, v := range convDirs {
		if !v.IsDir() {
			continue
		}

		filePath := filepath.Join(inboxPath, v.Name(), "message_1.json")
		conv, err := conversation.Parse(filePath)
		if err != nil {
			utils.PrintError("error: could not parse JSON file: %s", err)
		}

		convs = append(convs, conv)
	}
}
