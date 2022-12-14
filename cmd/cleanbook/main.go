package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/samsapti/CleanBook/internal/utils"
	"github.com/samsapti/CleanBook/pkg/conversation"
	"github.com/samsapti/CleanBook/pkg/user"
	web "github.com/samsapti/CleanBook/web/app"
)

const appTitle string = "CleanBook"

var (
	basePath *string = flag.String("path", "", "Path to the directory containing your Facebook data (required)")
	port     *int    = flag.Int("port", 8080, "Port to listen on")

	convs  map[string]*conversation.Conversation
	fbUser *user.Profile
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Quit if -path was not specified
	if len(*basePath) == 0 {
		flag.Usage()
		utils.PrintFatal("\nerror: -path must be specified")
	}

	// convs is a map from conversation path to Conversation struct
	convs = make(map[string]*conversation.Conversation)

	// Get conversation dirs
	messagesPath := filepath.Join(*basePath, "messages")
	inboxPath := filepath.Join(messagesPath, "inbox")
	convDirs, err := os.ReadDir(inboxPath)
	if err != nil {
		utils.PrintFatal("error: %s", err)
	}

	for _, v := range convDirs {
		if !v.IsDir() {
			continue
		}

		filePath := filepath.Join(inboxPath, v.Name(), "message_1.json")
		conv, err := conversation.Parse(filePath)
		if err != nil {
			utils.PrintError("error: %s", err)
		}

		convs[conv.Path] = conv
	}

	// Get user information
	profilePath := filepath.Join(*basePath, "profile_information", "profile_information.json")
	fbUser, err = user.Parse(profilePath)
	if err != nil {
		utils.PrintFatal("error: %s", err)
	}

	web.Serve(&web.RuntimeData{
		AppTitle: appTitle,
		User:     fbUser,
		Convs:    convs,
		Port:     *port,
	})
}
