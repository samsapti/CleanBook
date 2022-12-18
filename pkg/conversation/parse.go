package conversation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/imdario/mergo"
	"github.com/samsapti/CleanBook/internal/utils"
)

type messagesTransformer struct{}

func (t messagesTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf([]*Message{}) {
		return func(dst, src reflect.Value) error {
			if dst.CanSet() {
				dst.Set(reflect.AppendSlice(dst, src))
			}

			return nil
		}
	}

	return nil
}

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

// Parse takes a path to a conversation folder, and loads the data into
// a Conversation struct. It returns a pointer to the Conversation
// struct.
func Parse(convPath string) (*Conversation, error) {
	var conv Conversation

	// Get directory contents
	dir, err := os.ReadDir(convPath)
	if err != nil {
		return nil, err
	}

	// Read message files
	for _, f := range dir {
		if !strings.HasPrefix(f.Name(), "message_") {
			continue
		}

		var convPart Conversation

		filePath := filepath.Join(convPath, f.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(file).Decode(&convPart)
		if err != nil {
			return nil, err
		}

		// Merge structs, only appending Messages slice field
		mergo.Merge(&conv, convPart, mergo.WithTransformers(messagesTransformer{}))
	}

	// Make conv.Path relative to the path given in -path
	conv.Path = filepath.Join("messages", conv.Path)

	// Fix deleted users
	if len(conv.Participants) == 1 {
		conv.Participants = append([]*Participant{{Name: "Facebook user"}}, conv.Participants...)

		if len(conv.Title) == 0 {
			conv.Title = "Facebook user"
		}
	}

	// Reverse conv.Messages slice
	for i, j := 0, len(conv.Messages)-1; i < j; i, j = i+1, j-1 {
		conv.Messages[i], conv.Messages[j] = conv.Messages[j], conv.Messages[i]
	}

	// Fix encoding in strings
	fixEncoding(&conv)

	return &conv, nil
}
