package web

import (
	"github.com/samsapti/CleanMessages/pkg/conversation"
	"github.com/samsapti/CleanMessages/pkg/user"
)

// PageData holds the data that should be used for rendering the web
// application templates.
type PageData struct {
	AppTitle     string
	PageTitle    string
	User         *user.Profile
	Conversation *conversation.Conversation
}

// RuntimeData holds data that is used on runtime.
type RuntimeData struct {
	AppTitle string
	User     *user.Profile
	Convs    []*conversation.Conversation
	Port     int
}
