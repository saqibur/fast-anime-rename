package organizer

import (
	"fmt"
	"os"
)

func PreviewChanges(filenames []string, fileExtensions []string, properNames []string) {
	idx := 0
	for _, value := range filenames {
		outputFilename := fmt.Sprintf("%s%s", properNames[idx], value)
		fmt.Println(outputFilename)
		idx++
	}
}

func RenameFiles(path string, filenames []string, fileExtensions []string, dst []string) {
	for idx, filename := range filenames {
		outputFilename := fmt.Sprintf("%s%s", dst[idx], fileExtensions[idx])
		originalName := path + "/" + fmt.Sprintf("%s%s", filename, fileExtensions[idx])
		newName := path + "/FAR/" + outputFilename
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
