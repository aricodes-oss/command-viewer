package main

import "fmt"

func main() {
	logger.Info("Booting...")

	bots, err := AvailableBots()
	if err != nil {
		panic(err)
	}

	numBots := len((*bots).ToSlice())

	logger.Info(fmt.Sprintf(
		"Loaded and registered %v bots",
		numBots,
	))

	ListenAndServe(bots)
}
