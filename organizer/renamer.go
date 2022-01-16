package organizer

import (
	"fmt"
	"os"
	"regexp"
)

func PreviewChanges(filenames []string, fileExtensions []string, episodeNames []string) {
	invalid_characters := regexp.MustCompile(`[\/\\:\"?*|<>]`)

	for idx, filename := range filenames {
		outputFilename := fmt.Sprintf("%s => %s%s", filename, invalid_characters.ReplaceAllString(episodeNames[idx], ""), fileExtensions[idx])
		fmt.Println(outputFilename)
		idx++
	}
}

func RenameFiles(path string, filenames []string, fileExtensions []string, dst []string) {
	invalid_characters := regexp.MustCompile(`[\/\\:\"?*|<>]`)

	for idx, filename := range filenames {
		outputFilename := fmt.Sprintf("%s%s", dst[idx], fileExtensions[idx])
		originalName := path + "/" + fmt.Sprintf("%s%s", filename, fileExtensions[idx])
		newName := path + "/FAR/" + invalid_characters.ReplaceAllString(outputFilename, "")
		fmt.Println(originalName)
		fmt.Println(newName)
		idx++
		os.Mkdir(path+"/FAR", 0755)
		// Comment out the following line when testing things out.
		var err error = os.Rename(originalName, newName)
		if err != nil {
			panic(err)
		}
	}
}
