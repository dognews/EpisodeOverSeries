package main

import (
	"EpisodeOverSeries/internal/project"
	"errors"
	"flag"
	"os"
	"slices"
	"strconv"
)

func FetchSeriesId(command string) (int, error) {
	index := slices.Index(os.Args, command) + 1

	if index >= len(os.Args) {
		return -1, errors.New("series id not provided")
	}

	seriesId, err := strconv.Atoi(os.Args[index])
	if err != nil {
		return -1, err
	}

	return seriesId, nil
}

func ParseFlags() (string, string) {
	serverIp := flag.String("ip", "NONE", "The IP of the Emby server")
	apiKey := flag.String("apikey", "NONE", "The api key of the Emby server")
	flag.Parse()

	return *serverIp, *apiKey
}

func ValidateCredentials(serverIp string, apiKey string) error {
	project.SetEmbyCredentials(serverIp, apiKey)
	err := project.ValidateEmbyCredentials()
	if serverIp == "NONE" || apiKey == "NONE" || err != nil {
		return errors.New("invalid credentials provided")
	}

	return nil
}
