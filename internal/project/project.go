package project

import (
	"EpisodeOverSeries/internal/emby"
	"fmt"
)

func ApplyToSeries(seriesId int) error {
	err := emby.SetSeriesThumbnails(seriesId)
	if err != nil {
		return err
	}

	return nil
}

func RemoveFromSeries(seriesId int) error {
	err := emby.RemoveSeriesThumbnails(seriesId)
	if err != nil {
		return err
	}

	return nil
}

func ApplyToAllSeries() {
	series := emby.GetAllSeries()
	quantity := len(series)

	for index, show := range series {
		err := emby.SetSeriesThumbnails(show)
		if err != nil {
			fmt.Printf("FAILED: Series refused to update and has been skipped (%d/%d)!\n", index+1, quantity)
		} else {
			fmt.Printf("Series updated (%d/%d)!\n", index+1, quantity)
		}
	}
}

func RemoveFromAllSeries() {
	series := emby.GetAllSeries()
	quantity := len(series)

	for index, show := range series {
		err := emby.RemoveSeriesThumbnails(show)
		if err != nil {
			fmt.Printf("FAILED: Series refused to be deleted and has been skipped (%d/%d)!\n", index+1, quantity)
		} else {
			fmt.Printf("Series removed (%d/%d)!\n", index+1, quantity)
		}
	}
}

func ValidateEmbyCredentials() error {
	err := emby.IsServerOnline()
	if err != nil {
		return err
	}

	return nil
}

func SetEmbyCredentials(serverIp string, apiKey string) {
	emby.SetCredentials(serverIp, apiKey)
}
