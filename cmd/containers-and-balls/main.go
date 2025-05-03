package main

import (
	"fmt"

	"github.com/romanpitatelev/containers-and-balls/internal/reader"
	"github.com/romanpitatelev/containers-and-balls/internal/sorter"
	"github.com/rs/zerolog/log"
)

func main() {
	number, err := reader.ReadNumber()
	if err != nil {
		log.Panic().Err(err).Msg("invalid number")
	}

	containersSize, ballColorCounts, err := reader.ReadContainersBals(number)
	if err != nil {
		log.Panic().Err(err).Msg("failed to scan input data")
	}

	isPossible := sorter.SortAndCompareSlices(containersSize, ballColorCounts)

	if isPossible {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
