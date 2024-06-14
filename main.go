package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/robhorbury/radio_project/musictools"
)

func main() {
	api_key := loadApiKey()
	currentSong := musictools.GetNowPlaying("bbc6music", api_key)
	fmt.Println(currentSong)
}

func loadApiKey() string {
	file, err := os.ReadFile("config")
	if err != nil {
		panic(err)
	}

	key := strings.Split(strings.Split(string(file), "api_key=")[1], "\n")[0]
	return key
}
