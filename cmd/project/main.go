package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	serverIp, apiKey := ParseFlags()

	if slices.Contains(os.Args[1:], "help") {
		PrintHelp()
		return
	}

	err := ValidateCredentials(serverIp, apiKey)
	if err != nil {
		fmt.Println("invalid credentials provided")
		return
	}

	if slices.Contains(os.Args[1:], "applytoseries") {
		ApplyToSeries()
		return
	}

	if slices.Contains(os.Args[1:], "removefromseries") {
		RemoveFromSeries()
		return
	}

	if slices.Contains(os.Args[1:], "applytoall") {
		ApplyToAllSeries()
		return
	}

	if slices.Contains(os.Args[1:], "removefromall") {
		RemoveFromAllSeries()
		return
	}

	UnrecognizedCommand()
}
