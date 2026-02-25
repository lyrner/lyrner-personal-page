package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
)

// ---------------------------------------------------------------------------
// Embedded assets — templates and static files are baked into the binary,
// so the Docker image only needs the single compiled binary.
// ---------------------------------------------------------------------------

//go:embed templates
var templateFiles embed.FS

//go:embed static
var staticFiles embed.FS

// ---------------------------------------------------------------------------
// Data
// ---------------------------------------------------------------------------

// Project represents a subdomain project shown on the landing page.
type Project struct {
	Name        string
	Description string
	Subdomain   string
}

// projects is the list displayed on the landing page.
// Each entry links to https://<Subdomain>.lyrner.se
var projects = []Project{
	{
		Name:        "Project One",
		Description: "A short description of what this project does.",
		Subdomain:   "project-one",
	},
	{
		Name:        "Project Two",
		Description: "Another project living on its own subdomain.",
		Subdomain:   "project-two",
	},
}

// ---------------------------------------------------------------------------
// Templates
// ---------------------------------------------------------------------------

var tmpl = template.Must(template.ParseFS(templateFiles, "templates/*.html"))

// ---------------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------------

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "index.html", projects); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		log.Printf("template error: %v", err)
	}
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("static fs: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
