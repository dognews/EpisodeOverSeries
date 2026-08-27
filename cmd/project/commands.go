package main

import (
	"EpisodeOverSeries/internal/project"
	"fmt"
)

func PrintHelp() {
	fmt.Println("Usage: project [command]")
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
