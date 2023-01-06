package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/samsapti/CleanBook/internal/utils"
	"github.com/samsapti/CleanBook/pkg/conversation"
	"github.com/samsapti/CleanBook/pkg/user"
	"github.com/samsapti/CleanBook/web"
)

const appTitle string = "CleanBook"

var (
	basePath *string = flag.String("path", "", "Path to the directory containing your Facebook data (required)")
	port     *int    = flag.Int("port", 8080, "Port to listen on")
	verbose  *bool   = flag.Bool("verbose", false, "Enable verbose output")
	convs            = make(map[string]*conversation.Conversation)
	fbUser   *user.Profile
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Quit if -path was not specified
	if len(*basePath) == 0 {
		flag.Usage()
		utils.PrintFatal("\nerror: -path must be specified")
	}

	// Get conversation dirs
	utils.PrintVerbose(*verbose, "Locating conversation directories")
	messagesPath := filepath.Join(*basePath, "messages")
	inboxPath := filepath.Join(messagesPath, "inbox")
	convDirs, err := os.ReadDir(inboxPath)
	if err != nil {
		utils.PrintFatal("error: %s", err)
	}

	// Read conversation files
	utils.PrintVerbose(*verbose, "Parsing conversations")
	for _, v := range convDirs {
		if !v.IsDir() {
			continue
		}

		convPath := filepath.Join(inboxPath, v.Name())
		conv, err := conversation.Parse(convPath)
		if err != nil {
			utils.PrintError("error: %s", err)
		}

		// Map v.Name() to *Conversation
		convs[v.Name()] = conv
	}

	// Get profile information
	utils.PrintVerbose(*verbose, "Gathering profile information")
	profilePath := filepath.Join(*basePath, "profile_information", "profile_information.json")
	fbUser, err = user.Parse(profilePath)
	if err != nil {
		utils.PrintFatal("error: %s", err)
	}

	// Serve the application
	utils.PrintVerbose(*verbose, "Serving application with gathered data")
	web.Serve(&web.RuntimeData{
		AppTitle: appTitle,
		User:     fbUser,
		Convs:    convs,
		Port:     *port,
		BasePath: *basePath,
		Verbose:  *verbose,
	})
}
