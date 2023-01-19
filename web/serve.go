package web

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/samsapti/CleanBook/internal/utils"
	"github.com/samsapti/CleanBook/pkg/conversation"
)

//go:embed templates/*
var embedFS embed.FS
var rd *RuntimeData

func sanitize(filename string) string {
	sanitized := strings.ReplaceAll(filename, "\n", "")
	sanitized = strings.ReplaceAll(sanitized, "\r", "")
	return sanitized
}

func parseTemplates(filenames ...string) (*template.Template, error) {
	var tmplFiles []string
	tmplDir := "templates"
	tmplFiles = append(tmplFiles, filepath.Join(tmplDir, "partials", "navbar.html"))
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
			default:
				return base + "none"
			}
		},
		"conversationClass": func(t string) string {
			base := "conversation-type-"

			switch t {
			case conversation.ConversationRegular:
				return base + "regular"
			case conversation.ConversationRegularGroup:
				return base + "group"
			default:
				return base + "none"
			}
		},
		"isGroup": func(t string) bool {
			return t == conversation.ConversationRegularGroup
		},
		"nl2br": func(s string) template.HTML {
			r := strings.ReplaceAll(template.HTMLEscapeString(s), "\n", "<br>")
			return template.HTML(r)
		},
	}

	return template.New(tmplName).Funcs(funcMap).ParseFS(embedFS, tmplFiles...)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmplFile := "index.html"

	utils.PrintVerbose(rd.Verbose, "Parsing template %s", tmplFile)
	tmpl, err := parseTemplates(tmplFile)
	if err != nil {
		utils.PrintError("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.PrintVerbose(rd.Verbose, "Executing template %s", tmplFile)
	tmpl.Execute(w, &PageData{
		AppTitle:  rd.AppTitle,
		PageTitle: "Welcome",
		User:      rd.User,
		Convs:     rd.Convs,
	})
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	tmplFile := "messages.html"

	utils.PrintVerbose(rd.Verbose, "Parsing template %s", tmplFile)
	tmpl, err := parseTemplates(tmplFile)
	if err != nil {
		utils.PrintError("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.PrintVerbose(rd.Verbose, "Executing template %s", tmplFile)
	tmpl.Execute(w, &PageData{
		AppTitle:  rd.AppTitle,
		PageTitle: "Messages",
		User:      rd.User,
		Convs:     rd.Convs,
	})
}

func handleConv(w http.ResponseWriter, r *http.Request) {
	pageTitle := "Messages"
	convID := chi.URLParam(r, "convID")
	conv := rd.Convs[convID]
	if conv != nil {
		pageTitle += " - " + conv.Title
	}

	tmplFile := filepath.Join("partials", "conversation.html")

	utils.PrintVerbose(rd.Verbose, "Parsing template %s", tmplFile)
	tmpl, err := parseTemplates(tmplFile, "messages.html")
	if err != nil {
		utils.PrintError("error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.PrintVerbose(rd.Verbose, "Executing template %s", tmplFile)
	tmpl.Execute(w, &PageData{
		AppTitle:  rd.AppTitle,
		PageTitle: pageTitle,
		User:      rd.User,
		Convs:     rd.Convs,
		ConvID:    convID,
	})
}

func handleConvImage(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	filePath := filepath.Join(rd.BasePath, "messages", "photos", filename)

	utils.PrintVerbose(rd.Verbose, "Reading image data from %s", sanitize(filePath))
	imgData, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
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
	filePath := filepath.Join(rd.BasePath, "messages", "inbox", convID, fileType, filename)

	utils.PrintVerbose(rd.Verbose, "Reading file %s", sanitize(filePath))
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
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
	filePath := filepath.Join(rd.BasePath, "messages", "stickers_used", filename)

	utils.PrintVerbose(rd.Verbose, "Reading sticker data from %s", sanitize(filePath))
	stickerData, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
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
func Serve(data *RuntimeData) {
	rd = data

	// Prepare router
	utils.PrintVerbose(rd.Verbose, "Preparing router with middlewares")
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://localhost:%d", rd.Port)},
		AllowedMethods:   []string{"GET"},
		AllowCredentials: false,
		Debug:            rd.Verbose,
	}))

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
