package jikan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const jikanRequestEndpoint string = "https://api.jikan.moe/v3/"

// This seemed like a good trade between quantity and quality of search results.
const searchResultLimit uint = 7

// Need to change this to a uint.
type MalId int

type AnimeSearchResults struct {
	Results []struct {
		MalID        int    `json:"mal_id"`
		URL          string `json:"url"`
		Title        string `json:"title"`
		EpisodeCount int    `json:"episodes"`
	} `json:"results"`
}

type AnimeEpisodes struct {
	EpisodesLastPage int `json:"episodes_last_page"`
	Episodes         []struct {
		EpisodeID int    `json:"episode_id"`
		Title     string `json:"title"`
	} `json:"episodes"`
}

type Anime struct {
	MalID        int    `json:"mal_id"`
	Title        string `json:"title"`
	TitleEnglish string `json:"title_english"`
}

func Search(searchTerms string) AnimeSearchResults {
	fmt.Println("Searching")
	searchEndpoint := jikanRequestEndpoint + "search/anime"
	animeSearchQueryValues := url.Values{}
	animeSearchQueryValues.Add(
		"limit",
		strconv.FormatUint(uint64(searchResultLimit), 10 /* base */),
	)
	animeSearchQueryValues.Add("q", searchTerms)

	// This is async, need to figure out how to deal with this properly. And not
	// like an idiot.
	resp, err := http.Get((searchEndpoint + "?" + animeSearchQueryValues.Encode()))

	if err != nil {
		fmt.Println(err)
		return AnimeSearchResults{}
	}

	searchResult := &AnimeSearchResults{}
	var _ error = json.NewDecoder(resp.Body).Decode(searchResult)

	return *searchResult
}

func getAnimeTitle(malId MalId) string {
	resp, _ := http.Get(("https://api.jikan.moe/v3/anime" + "/" + strconv.Itoa(int(malId))))
	searchResult := &Anime{}
	var _ error = json.NewDecoder(resp.Body).Decode(searchResult)
	return searchResult.TitleEnglish
}

func RetrieveEpisodeList(malId MalId) []string {
	resp, _ := http.Get(("https://api.jikan.moe/v3/anime" + "/" + strconv.Itoa(int(malId)) + "/episodes" + "/1"))
	searchResult := &AnimeEpisodes{}
	var _ error = json.NewDecoder(resp.Body).Decode(searchResult)

	// Need to check for episode's last page, iterate through all the pages, respecting
	// the rate limit, and then appending everything to a list.

	// TODO: You need to let the user choose between title and english title.
	// HACK: this is a bad way to do it.
	showEnglishTitle := getAnimeTitle(malId)

	var episodes []string
	for _, episode := range searchResult.Episodes {
		// Needs to be able to handle episode counts > 100. needs some changes obviously.
		formattedName := fmt.Sprintf("%s %02d - %s", showEnglishTitle, episode.EpisodeID, episode.Title)
		episodes = append(episodes, formattedName)
	}

	return episodes
}
