package musictools

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type recentlyPlayedResponse struct {
	// Structure to handle recently played requests
	Recenttracks struct {
		Track []struct {
			Artist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type Song struct {
	Name         string
	Artist       string
	URL          string
	RadioStation string
}

func getRecentlyPlayedJson(user, api_key string) []byte {
	// Get the recently played song list from a chosen user
	resp, err := http.Get(fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%v&api_key=%v&format=json", user, api_key))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func parseRecentlyPlayed(response []byte) recentlyPlayedResponse {

	var parsed recentlyPlayedResponse
	err := json.Unmarshal(response, &parsed)
	if err != nil {
		log.Fatalln(err)
	}
	return parsed
}

func GetNowPlaying(user, api_key string) Song {

	recentlyPlayedJson := getRecentlyPlayedJson(user, api_key)
	parsedRecentlyPlayed := parseRecentlyPlayed(recentlyPlayedJson)
	nowPlaying := parsedRecentlyPlayed.Recenttracks.Track[0]

	return Song{
		Name:         nowPlaying.Name,
		Artist:       nowPlaying.Artist.Text,
		URL:          nowPlaying.URL,
		RadioStation: user}

}
