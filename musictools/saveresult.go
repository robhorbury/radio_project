package musictools

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const catalog string = "./../radio_project_database"

func RecordSongPlayed(playedSong Song) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%v/radio.db", catalog))
	if err != nil {
		panic(err)
	}

	if playedSong.Name != getLastSongSaved() {
		dts := time.Now().String()

		sqlString := fmt.Sprintf(`
	INSERT INTO song_playing_raw VALUES
	(null, "%v", "%v", "%v", "%v", "%v")
	`, playedSong.Name,
			playedSong.Artist,
			playedSong.URL,
			playedSong.RadioStation,
			dts)

		_, err = db.Exec(sqlString)
		if err != nil {
			panic(err)
		}
	}
}

func getLastSongSaved() string {

	db, err := sql.Open("sqlite3", fmt.Sprintf("%v/radio.db", catalog))
	if err != nil {
		panic(err)
	}

	sqlString := `
	SELECT song_name FROM song_playing_raw
	ORDER BY id DESC
	LIMIT 1
	`
	var songName string

	err = db.QueryRow(sqlString).Scan(&songName)
	if err != nil && !(errors.Is(err, sql.ErrNoRows)) {
		panic(err)
	}
	return songName
}

func InitDatabase() {
	createSongPlayingRaw()
}

func createSongPlayingRaw() {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%v/radio.db", catalog))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS song_playing_raw (
		id INTEGER NOT NULL PRIMARY KEY,
		song_name TEXT,
		artist TEXT,
		url TEXT,
		radio_station TEXT,
		dts TEXT
		);`)

	if err != nil {
		panic(err)
	}

}
