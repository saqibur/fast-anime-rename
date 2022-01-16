package browser

import (
	"bufio"
	"fast-anime-rename/jikan"
	"fmt"
	"os"
	"strconv"
)

func SearchAnime() jikan.MalId {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Search for: ")
	input, err := reader.ReadString('\n')
	fmt.Println("-------------------------")

	if err != nil {
		panic(err)
	}

	searchResults := jikan.Search(input)

	if len(searchResults.Results) != 0 {
		fmt.Println("Selection - EpisodeCount - Title - URL")
		for idx, anime := range searchResults.Results {
			output := fmt.Sprintf("%d\t%d\t%s - %s", idx, anime.EpisodeCount, anime.Title, anime.URL)
			fmt.Println(output)
		}
	}

	fmt.Println("Select the correct anime")

	animeSelection, _ := reader.ReadString('\n')
	selectedIndex, _ := strconv.Atoi(animeSelection)

	selectedAnime := searchResults.Results[selectedIndex]
	return jikan.MalId(selectedAnime.MalID)
}
