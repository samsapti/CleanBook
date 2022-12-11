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

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &conv); err != nil {
		return nil, err
	}

	// Make conv.Path relative to the path given in -d
	conv.Path = filepath.Join("messages", conv.Path)

	return &conv, nil
}
