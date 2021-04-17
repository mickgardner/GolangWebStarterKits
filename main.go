package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var staticFiles embed.FS

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	// This is inefficient - it reads the templates from the filesystem every
	// time. This makes it much easier to develop though, so we can edit our
	// templates and the changes will be reflected without having to restart
	// the app.
	t, err := template.ParseFS(templates, "templates/*.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err.Error()), 500)
		return
	}

	err = t.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %s", err.Error()), 500)
		return
	}
}

func logRequest(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)

		f(w, r)
	})
}

func (a App) logStatic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("path: %s", r.URL.Path)

		h.ServeHTTP(w, r)
	})
}

type App struct {
	Port string
}

func (a App) Start() {
	log.Printf("serving static assets")

	http.Handle("/static/", a.logStatic(http.FileServer(http.FS(staticFiles))))

	http.Handle("/", logRequest(a.index))
	addr := fmt.Sprintf(":%s", a.Port)
	log.Printf("Starting app on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func (a App) index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", struct {
		Name string
	}{
		Name: "Sonic The Hedgehog",
	})
}

func env(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func main() {
	server := App{
		Port: env("PORT", "8080"),
	}
	server.Start()
}
