package organizer

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadDirectory(pathToDirectory string) ([]string, []string) {
	files, err := os.ReadDir(strings.TrimSpace(pathToDirectory))

	if err != nil {
		log.Fatal(err)
	}

	var filenames []string
	var fileExtensions []string

	for _, file := range files {
		filename := file.Name()
		filenames = append(filenames, strings.TrimSuffix(filename, filepath.Ext(filename)))
		fileExtensions = append(fileExtensions, filepath.Ext(filename))
	}

	return filenames, fileExtensions
}
