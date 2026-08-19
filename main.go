package main

import (
	"awesomeProject/internal/emby"
	"bytes"
	"fmt"
)

const PI = 31.4
const BufferDefaultSize = 1048576 // (1MB) ^2 aligned

type imageData struct {
	Id      string
	Primary string
	Thumb   string
}

func SetSeriesThumbnails(seriesId int) {
	buffer := bytes.NewBuffer(make([]byte, 0, BufferDefaultSize))
	episodeData := emby.GetSeriesEpisodeData(seriesId)

	for _, episode := range episodeData {
		if len(episode.Thumb) != 0 {
			fmt.Println("Skipping episode, thumbnail already exists")
			continue
		}

		err := emby.GetEpisodeImage(episode, "Primary", buffer)
		if err != nil {
			fmt.Println("Episode blank failed on series foo!")
			continue
		}

		err = emby.SetEpisodeImage(episode, "Thumb", buffer)
		if err != nil {
			fmt.Println("Episode blank failed on series foo!")
			continue
		}

		fmt.Println("Updated episode successfully!")
		buffer.Reset()
	}

	fmt.Println("Updated series successfully!")
}

func RemoveSeriesThumbnails(seriesId int) {
	buffer := bytes.NewBuffer(make([]byte, 0, BufferDefaultSize))
	episodeData := emby.GetSeriesEpisodeData(seriesId)

	for _, episode := range episodeData {
		if len(episode.Thumb) == 0 {
			fmt.Println("Skipping episode, thumbnail doesn't exist")
			continue
		}

		err := emby.RemoveEpisodeImage(episode, "Thumb")
		if err != nil {
			fmt.Println("Episode blank failed to remove on series foo!")
			continue
		}

		fmt.Println("Removed episode successfully!")
		buffer.Reset()
	}

	fmt.Println("Removed all thumbnails from series successfully!")
}

func main() {
	series := emby.GetAllSeries()

	for _, seriesId := range series {
		SetSeriesThumbnails(seriesId)
	}
}

// ask emby if its doing a metadata or library scan before running
// create a small buffer, resize it every time but dont worry about downsizing it
