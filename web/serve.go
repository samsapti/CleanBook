package web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samsapti/CleanBook/internal/utils"
	"github.com/samsapti/CleanBook/pkg/conversation"
	"github.com/samsapti/CleanBook/pkg/user"
)

var (
	//go:embed templates/*
	fs       embed.FS
	appTitle string
	fbUser   *user.Profile
	convs    map[string]*conversation.Conversation
	basePath string
)

func parseTemplates(filenames ...string) (*template.Template, error) {
	var tmplFiles []string
	tmplDir := "templates"
	tmplFiles = append(tmplFiles, filepath.Join(tmplDir, "layout.html"))

	// Append filenames
	for _, v := range filenames {
		tmplFiles = append(tmplFiles, filepath.Join(tmplDir, v))
	}

	tmplName := filenames[len(filenames)-1]
	funcMap := template.FuncMap{
		"base": func(path string) string {
			return filepath.Base(path)
		},
		"fromUnix": func(ts int64) string {
			return time.Unix(ts, 0).String()
		},
		"fromUnixMS": func(ts int64) string {
			return time.UnixMilli(ts).String()
		},
		"messageClass": func(t string) string {
			base := "message-type-"

			switch t {
			case conversation.MessageGeneric:
				return base + "generic"
			case conversation.MessageShare:
				return base + "share"
			case conversation.MessageSubscribe:
				return base + "subscribe"
			case conversation.MessageUnsubscribe:
				return base + "unsubscribe"
			case conversation.MessageCall:
				return base + "call"
			}

			return ""
		},
		"conversationClass": func(t string) string {
			base := "conversation-type-"

			switch t {
			case conversation.ConversationRegular:
				return base + "regular"
			case conversation.ConversationRegularGroup:
				return base + "group"
			}

			return ""
		},
		"isGroup": func(t string) bool {
			return t == conversation.ConversationRegularGroup
		},
		"nl2br": func(s string) template.HTML {
			r := strings.ReplaceAll(template.HTMLEscapeString(s), "\n", "<br>")
			return template.HTML(r)
		},
	}

	return template.New(tmplName).Funcs(funcMap).ParseFS(fs, tmplFiles...)
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

	convTmpl := filepath.Join("partials", "conversation.html")
	tmpl, err := parseTemplates(convTmpl, "messages.html")
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

func handleConvImage(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	filePath := filepath.Join(basePath, "messages", "photos", filename)

	imgData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			utils.PrintError("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Write(imgData)
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	convID := chi.URLParam(r, "convID")
	fileType := chi.URLParam(r, "fileType")
	filename := chi.URLParam(r, "filename")
	filePath := filepath.Join(basePath, "messages", "inbox", convID, fileType, filename)

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			utils.PrintError("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Write(fileData)
}

func handleSticker(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	filePath := filepath.Join(basePath, "messages", "stickers_used", filename)

	stickerData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			utils.PrintError("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Write(stickerData)
}

// Serve renders the web application and serves it on the port in
// rd.Port. rd is a RuntimeData struct.
func Serve(rd *RuntimeData) {
	appTitle = rd.AppTitle
	fbUser = rd.User
	convs = rd.Convs
	basePath = rd.BasePath

	// Prepare router
	utils.PrintVerbose(rd.Verbose, "Preparing router with middlewares")
	r := chi.NewRouter()
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)
	if rd.Verbose {
		r.Use(middleware.Logger)
	}

	// Setup routes
	utils.PrintVerbose(rd.Verbose, "Setting routes")
	r.Get("/", handleIndex)
	r.Get("/messages", handleMessages)
	r.Get("/messages/{convID}", handleConv)
	r.Get("/images/messages/photos/{filename}", handleConvImage)
	r.Get("/files/messages/inbox/{convID}/{fileType}/{filename}", handleFile)
	r.Get("/stickers/messages/stickers_used/{filename}", handleSticker)

	// Serve application
	utils.PrintInfo("Listening on http://localhost:%d", rd.Port)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", rd.Port), r); err != nil {
		utils.PrintFatal("error: %s", err)
	}
}
