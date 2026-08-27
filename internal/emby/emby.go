package emby

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"resty.dev/v3"
)

var client = resty.New()

func IsServerOnline() error {
	resp, err := client.R().SetTimeout(2*time.Second).SetQueryParams(map[string]string{
		"api_key": apiKey,
	}).SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("%s/emby/System/Info", serverIp))

	if err != nil {
		return err
	}

	if resp.IsStatusFailure() {
		return errors.New("server rejected credentials")
	}

	return nil
}

func GetAllSeries() []int {
	var series []int
	returnedItems := pageLength
	currentIndex := 0

	for returnedItems == pageLength {
		var content itemsRequest
		resp, err := client.R().SetQueryParams(map[string]string{
			"api_key":          apiKey,
			"IncludeItemTypes": "Series",
			"Fields":           "Id",
			"StartIndex":       strconv.Itoa(currentIndex),
			"SortBy":           "SortName",
			"SortOrder":        "Ascending",
			"Recursive":        "true",
			"Limit":            "50",
		}).SetHeader("Accept", "application/json").
			SetResult(&content).Get(
			fmt.Sprintf("%s/emby/Items", serverIp))

		if err != nil || resp.IsStatusFailure() {
			log.Fatal("Failed to fetch tv shows!")
		}

		for _, item := range content.Items {
			formattedId, _ := strconv.Atoi(item.Id)
			series = append(series, formattedId)
		}

		returnedItems = len(content.Items)
		currentIndex += pageLength
	}

	return series
}

func GetSeriesEpisodeData(seriesId int) ([]ImageData, error) {
	var episodeData []ImageData
	returnedItems := pageLength
	currentIndex := 0

	for returnedItems == pageLength {
		var content itemsRequest
		resp, err := client.R().SetQueryParams(map[string]string{
			"api_key":    apiKey,
			"ParentId":   strconv.Itoa(seriesId),
			"IsFolder":   "false",
			"Fields":     "Id",
			"StartIndex": strconv.Itoa(currentIndex),
			"Recursive":  "true",
			"Limit":      "50",
		}).
			SetResult(&content).Get(
			fmt.Sprintf("%s/emby/Items", serverIp))

		if err != nil || resp.IsStatusFailure() {
			return []ImageData{}, err
		}

		for _, item := range content.Items {
			episodeData = append(episodeData, ImageData{item.Id, item.ImageTags["Primary"], item.ImageTags["Thumb"]})
		}

		returnedItems = len(content.Items)
		currentIndex += pageLength
	}

	if len(episodeData) == 0 {
		return episodeData, errors.New("no episode data found")
	}

	return episodeData, nil
}

func GetEpisodeImage(episodeImageData ImageData, imageType string, imageBuffer *bytes.Buffer) error {
	serverIp := fmt.Sprintf("%s/emby/Items/%s/Images/%s", serverIp, episodeImageData.Id, imageType)
	resp, err := client.R().SetQueryParams(map[string]string{
		"api_key": apiKey,
	}).Get(serverIp)

	defer resp.Body.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(imageBuffer, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func SetEpisodeImage(episodeImageData ImageData, imageType string, imageBuffer *bytes.Buffer) error {
	serverip := fmt.Sprintf("%s/emby/Items/%s/Images/%s", serverIp, episodeImageData.Id, imageType)
	_, err := client.R().
		SetQueryParam("api_key", apiKey).
		SetHeader("Content-Type", "image/jpeg").
		SetBody(base64.StdEncoding.EncodeToString(imageBuffer.Bytes())).
		Post(serverip)

	return err
}

func RemoveEpisodeImage(episodeImageData ImageData, imageType string) error {
	serverip := fmt.Sprintf("%s/emby/Items/%s/Images/%s/Delete", serverIp, episodeImageData.Id, imageType)
	_, err := client.R().SetQueryParam("api_key", apiKey).Post(serverip)

	return err
}

func SetCredentials(ip string, api string) {
	serverIp = ip
	apiKey = api
}
