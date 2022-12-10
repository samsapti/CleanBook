package conversation

// Participant represents a person participating in a
// conversation.
type Participant struct {
	Name string `json:"name"`
}

// JoinMode contains information on how a group conversation
// might be joined by others, and optionally includes a link
// to join the group conversation.
type JoinMode struct {
	Mode int    `json:"mode"`
	Link string `json:"link"`
}

// Reaction represents a reaction on a message.
type Reaction struct {
	Emoji string `json:"reaction"`
	Actor string `json:"actor"`
}

// Share represents a shared link.
type Share struct {
	Link string `json:"link"`
}

// Sticker represents a sticker message. Sticker.Path
// is the local path to the sticker's image file, relative
// to the path of the data directory.
type Sticker struct {
	Path string `json:"uri"`
}

// File represents files sent in a message. This includes
// photo messages. File.Path is the local path to the file,
// relative to the path of the data directory.
type File struct {
	Path      string `json:"uri"`
	TimeStamp uint64 `json:"creation_timestamp"`
}

// Message represents a message in a conversation. Some
// values are mutually exclusive, such as Message.Sticker
// and Message.Content.
type Message struct {
	SenderName   string     `json:"sender_name"`
	TimeStamp    uint64     `json:"timestamp_ms"`
	Content      string     `json:"content"`
	Files        []File     `json:"files"`
	Photos       []File     `json:"photos"`
	Share        Share      `json:"share"`
	Sticker      Sticker    `json:"sticker"`
	Reactions    []Reaction `json:"reactions"`
	CallDuration int        `json:"call_duration"`
	Type         string     `json:"type"`
	Unsent       bool       `json:"is_unsent"`
}

// Conversation represents a conversation with one or more
// Facebook user(s). It matches the data in the JSON files
// with a file path that looks like this:
// messages/inbox/{some_conversation}/message_1.json
type Conversation struct {
	Participants     []Participant `json:"participants"`
	Messages         []Message     `json:"messages"`
	Title            string        `json:"title"`
	StillParticipant bool          `json:"is_still_participant"`
	Type             string        `json:"thread_type"`
	Path             string        `json:"thread_path"`
	JoinMode         JoinMode      `json:"joinable_mode"`
}