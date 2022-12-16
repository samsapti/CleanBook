package web

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samsapti/CleanBook/internal/utils"
	"github.com/samsapti/CleanBook/pkg/conversation"
	"github.com/samsapti/CleanBook/pkg/user"
)

var (
	appTitle string
	fbUser   *user.Profile
	convs    map[string]*conversation.Conversation
)

func parseTemplates(fileName string) (*template.Template, error) {
	tmplDir := filepath.Join("web", "templates")
	layout := filepath.Join(tmplDir, "layout.html")
	tmpl := filepath.Join(tmplDir, fileName)

	return template.New(fileName).ParseFiles(layout, tmpl)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parseTemplates("index.html")
	if err != nil {
		utils.PrintError("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, &PageData{
		AppTitle:  appTitle,
		PageTitle: "Welcome",
		User:      fbUser,
		Convs:     convs,
	})
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	tmpl, err := parseTemplates("messages.html")
	if err != nil {
		utils.PrintError("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, &PageData{
		AppTitle:  appTitle,
		PageTitle: "Messages",
		User:      fbUser,
		Convs:     convs,
	})
}

func handleConv(w http.ResponseWriter, r *http.Request) {
	pageTitle := "Messages"
	convID := chi.URLParam(r, "convID")
	conv := convs[convID]
	if conv != nil {
		pageTitle += " - " + conv.Title
	}

	tmpl, err := parseTemplates("conversation.html")
	if err != nil {
		utils.PrintError("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, &PageData{
		AppTitle:  appTitle,
		PageTitle: pageTitle,
		User:      fbUser,
		Convs:     convs,
		ConvID:    chi.URLParam(r, "convID"),
	})
}

// Serve renders the web application and serves it on the port in
// rd.Port. rd is a RuntimeData struct.
func Serve(rd *RuntimeData) {
	appTitle = rd.AppTitle
	fbUser = rd.User
	convs = rd.Convs

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Setup routes
	r.Get("/", handleIndex)
	r.Get("/messages", handleMessages)
	r.Get("/messages/{convID}", handleConv)

	// Serve the application
	utils.PrintInfo("Listening on http://localhost:%d", rd.Port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", rd.Port), r); err != nil {
		utils.PrintFatal("error: %s", err)
	}
}
