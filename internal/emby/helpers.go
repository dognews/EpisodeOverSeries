package emby

import (
	"bytes"
)

func SetSeriesThumbnails(seriesId int) error {
	buffer := bytes.NewBuffer(make([]byte, 0, BufferDefaultSize))
	episodeData, err := GetSeriesEpisodeData(seriesId)
	if err != nil {
		return err
	}

	for _, episode := range episodeData {
		if len(episode.Thumb) != 0 {
			continue
		}

		err := GetEpisodeImage(episode, "Primary", buffer)
		if err != nil {
			continue
		}

		err = SetEpisodeImage(episode, "Thumb", buffer)
		if err != nil {
			continue
		}

		buffer.Reset()
	}

	return nil
}

func RemoveSeriesThumbnails(seriesId int) error {
	buffer := bytes.NewBuffer(make([]byte, 0, BufferDefaultSize))
	episodeData, err := GetSeriesEpisodeData(seriesId)
	if err != nil {
		return err
	}

	for _, episode := range episodeData {
		if len(episode.Thumb) == 0 {
			continue
		}

		err := RemoveEpisodeImage(episode, "Thumb")
		if err != nil {
			continue
		}

		buffer.Reset()
	}

	return nil
}
