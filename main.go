package main

import (
	"os"
	"strings"
	"time"

	"github.com/robhorbury/radio_project/musictools"
)

func main() {
	counter := 0
	musictools.InitDatabase()
	api_key := loadApiKey()
	for counter < 15 {
		currentSong := musictools.GetNowPlaying("bbc6music", api_key)
		musictools.RecordSongPlayed(currentSong)
		time.Sleep(60 * time.Second)
		counter += 1
	}

}

func loadApiKey() string {
	file, err := os.ReadFile("config")
	if err != nil {
		panic(err)
	}

	key := strings.Split(strings.Split(string(file), "api_key=")[1], "\n")[0]
	return key
}
