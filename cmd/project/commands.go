package main

import (
	"EpisodeOverSeries/internal/project"
	"fmt"
)

func PrintHelp() {
	fmt.Println("Usage: project [command]")
	fmt.Println("```\nUsage:\n    me.exe [flags] [command]\nCommands:\n     applytoseries [seriesid]    " +
		"Apply episode thumbnails to series\n     removefromseries [seriesid] Remove episode thumbnails from" +
		" series\n     applytoall                  Apply episode thumbnails to every series\n     removefrom" +
		"all               Remove episode thumbnails from every series\n     help                        Sho" +
		"w cli commands\nFlags:\n    -ip [ip-address]  The ip address of the emby server  (REQUIRED)\n    -a" +
		"pikey [api-key] The api key of the emby server.    (REQUIRED)\n```")
}

func ApplyToSeries() {
	seriesId, err := FetchSeriesId("apply")
	if err != nil {
		fmt.Println("")
		return
	}

	err = project.ApplyToSeries(seriesId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func RemoveFromSeries() {
	seriesId, err := FetchSeriesId("remove")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = project.RemoveFromSeries(seriesId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func ApplyToAllSeries() {
	project.ApplyToAllSeries()
}

func RemoveFromAllSeries() {
	project.RemoveFromAllSeries()
}

func UnrecognizedCommand() {
	fmt.Println("No command provided: see correct usage <me.exe help>")
}
