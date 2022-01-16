package main

import (
	// "fast-anime-rename/browser"
	"bufio"
	"fast-anime-rename/browser"
	"fast-anime-rename/jikan"
	"fast-anime-rename/organizer"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to FAR.")
	fmt.Println("I'm not very smart, but I can make your life a tiny bit easier.")
	fmt.Println("Make a backup copy of your files before doing anything.")
	fmt.Println("---------------")

	animeId := browser.SearchAnime()
	episodes := jikan.RetrieveEpisodeList(animeId)

	fmt.Println("Paste the directory of your anime:")
	reader := bufio.NewReader(os.Stdin)
	directory, _ := reader.ReadString('\n')

	filenames, fileExtensions := organizer.ReadDirectory(directory)
	organizer.PreviewChanges(filenames, fileExtensions, episodes)

	fmt.Println("Start rename? (Y/N)")
	renameDecision, _ := reader.ReadString('\n')
	renameDecision = strings.TrimSpace(renameDecision)

	if renameDecision == "Y" || renameDecision == "y" {
		fmt.Println("Renaming files...")
		organizer.RenameFiles(strings.TrimSpace(directory), filenames, fileExtensions, episodes)
	} else {
		fmt.Println("Aborted.")
		fmt.Println("Closing.")
		duration := time.Duration(5) * time.Second
		time.Sleep(duration)
	}
}
