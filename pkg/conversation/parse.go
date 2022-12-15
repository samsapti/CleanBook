package conversation

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Parse takes a path to a JSON file containing a conversation, and
// loads the data into a Conversation struct. It returns a pointer to
// the Conversation struct.
func Parse(filePath string) (*Conversation, error) {
	var conv Conversation

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Read JSON data
	if err = json.NewDecoder(file).Decode(&conv); err != nil {
		return nil, err
	}

	// Make conv.Path relative to the path given in -path
	conv.Path = filepath.Join("messages", conv.Path)

	// Reverse conv.Messages slice
	for i, j := 0, len(conv.Messages)-1; i < j; i, j = i+1, j-1 {
		conv.Messages[i], conv.Messages[j] = conv.Messages[j], conv.Messages[i]
	}

	return &conv, nil
}
