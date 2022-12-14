package web

import (
	"github.com/samsapti/CleanBook/pkg/conversation"
	"github.com/samsapti/CleanBook/pkg/user"
)

// PageData holds the data that should be used for rendering the web
// application templates.
type PageData struct {
	AppTitle    string
	PageTitle   string
	User        *user.Profile
	Convs       map[string]*conversation.Conversation
	CurrentConv *conversation.Conversation
}

// RuntimeData holds data that is used on runtime.
type RuntimeData struct {
	AppTitle string
	User     *user.Profile
	Convs    map[string]*conversation.Conversation
	Port     int
}
