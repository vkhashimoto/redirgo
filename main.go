package main

import (
	"log"
	"net/http"
	"os"
	"redirgo/links"
)
func getLink(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	path := r.PathValue("link")
	log.Printf("The requested URL is: %s/%s\n", host, path)
	log.Printf("The requested link is: %s\n ", path)
	redirectTo, err := links.FindRedirection(host, path)
	if err != nil {
		//TODO: Show that the link was not found
		renderNotFound(w, r)
		return
	}
	w.Header().Add("Referer", "no-referrer")
	w.Header().Add("Location", redirectTo)
	w.WriteHeader(http.StatusTemporaryRedirect)
	return
}

func renderNotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func main() {
	addr := ":8080"
	linksFilePath := os.Getenv("LINKS_FILE")
	if linksFilePath == "" {
		log.Printf("No path provided by env var (LINKS_FILE)\n")
		linksFilePath = "./config/links.toml"
		log.Printf("Using default path (%s)\n", linksFilePath)
	}
	links.LoadLinks(linksFilePath)

	mux := http.NewServeMux()
	mux.HandleFunc("/{link}", getLink)
	mux.HandleFunc("/", root)

	log.Printf("Listening on `%s`", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
