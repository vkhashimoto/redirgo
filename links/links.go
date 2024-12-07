package links

import (
	"errors"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var links map[string]map[string]string

func LoadLinks(linksFilePath string) {
	if linksFilePath == "" {
		log.Fatal("Links file path cannot be empty")
	}
	links = make(map[string]map[string]string)
	if _, err := os.Stat(linksFilePath); err != nil {
		log.Fatalf("File `%s` not found", linksFilePath)
	}
	_, err := toml.DecodeFile(linksFilePath, &links)
	if err != nil {
		log.Fatalf("Error while decoding file: %s", err)
		os.Exit(1)
	}
}

func FindRedirection(host, path string) (string, error) {
	hostLinks, exists := links[host]
	if !exists {
		hostLinks, exists = links["*"]
		if !exists {
			return "", errors.New("Link not found")
		}
	}
	link, exists := hostLinks[path]
	if !exists {
		return "", errors.New("Link not found")
	}

	return link, nil
}
