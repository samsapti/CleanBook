package conversation

// Participant represents a Facebook user participating in a
// conversation.
type Participant struct {
	Name string `json:"name"`
}

// JoinMode contains information on how a group conversation might be
// joined by others, and optionally includes a link to join the group
// conversation.
type JoinMode struct {
	Mode int    `json:"mode"`
	Link string `json:"link"`
}

// Reaction represents a reaction to a message.
type Reaction struct {
	Emoji string `json:"reaction"`
	Actor string `json:"actor"`
}

// SharedMedia represents a shared media.
type SharedMedia struct {
	Link string `json:"link"`
}

// File represents files sent in a message. This includes photos and
// stickers. File.Path is the local path to the file, relative to the
// path of the data directory.
type File struct {
	Path      string `json:"uri"`
	TimeStamp int64  `json:"creation_timestamp"`
}

// Message represents a message in a conversation. Some values are
// mutually exclusive, such as Message.Sticker and Message.Content.
type Message struct {
	SenderName   string         `json:"sender_name"`
	TimeStampMS  int64          `json:"timestamp_ms"`
	Content      string         `json:"content"`
	Audio        []*File        `json:"audio_files"`
	Files        []*File        `json:"files"`
	Photos       []*File        `json:"photos"`
	Videos       []*File        `json:"videos"`
	Share        *SharedMedia   `json:"share"`
	Sticker      *File          `json:"sticker"`
	Reactions    []*Reaction    `json:"reactions"`
	CallDuration int64          `json:"call_duration"`
	Type         string         `json:"type"`
	Users        []*Participant `json:"users"`
	Unsent       bool           `json:"is_unsent"`
}

// Conversation represents a conversation with one or more Facebook
// user(s). It matches the data in the JSON files with a file path that
// looks like: messages/inbox/{convID}/message_{num}.json
type Conversation struct {
	Participants     []*Participant `json:"participants"`
	Messages         []*Message     `json:"messages"`
	Title            string         `json:"title"`
	StillParticipant bool           `json:"is_still_participant"`
	Type             string         `json:"thread_type"`
	Path             string         `json:"thread_path"`
	Image            *File          `json:"image"`
	JoinMode         *JoinMode      `json:"joinable_mode"`
}
