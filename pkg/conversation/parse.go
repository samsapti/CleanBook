package conversation

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/samsapti/CleanBook/internal/utils"
)

func fixEncoding(c *Conversation) {
	utils.FixEncoding(&c.Title)

	for _, p := range c.Participants {
		utils.FixEncoding(&p.Name)
	}

	for _, m := range c.Messages {
		utils.FixEncoding(&m.SenderName)
		utils.FixEncoding(&m.Content)

		for _, r := range m.Reactions {
			utils.FixEncoding(&r.Actor)
			utils.FixEncoding(&r.Emoji)
		}
	}
}

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

	// Fix encoding in strings
	fixEncoding(&conv)

	return &conv, nil
}
