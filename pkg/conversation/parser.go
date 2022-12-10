package conversation

import (
	"encoding/json"
	"os"

	"github.com/samsapti/CleanMessages/internal/utils"
)

// Parse takes a path to a JSON file containing a conversation, and
// loads the data into a Conversation struct. It returns a pointer to
// the Conversation struct.
func Parse(filePath string) (*Conversation, error) {
	var conv Conversation

	// Read JSON as raw data from file
	data, err := os.ReadFile(filePath)
	if err != nil {
		utils.PrintError("Error reading from %s: %s", filePath, err)
		return nil, err
	}

	// Unmarshall the JSON data
	if err = json.Unmarshal(data, &conv); err != nil {
		utils.PrintError("Error reading JSON data from %s: %s", filePath, err)
		return nil, err
	}

	return &conv, nil
}
