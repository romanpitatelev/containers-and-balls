package main

import (
	scannerservice "github.com/romanpitatelev/containers-and-balls/internal/scanner-service"
	sortingservice "github.com/romanpitatelev/containers-and-balls/internal/sorting-service"
	writerservice "github.com/romanpitatelev/containers-and-balls/internal/writer-service"
	"github.com/rs/zerolog/log"
)

func main() {
	scanner := scannerservice.New()
	writer := writerservice.New()

	containersSize, ballColorCounts, err := scanner.Scan()
	if err != nil {
		log.Panic().Err(err).Msg("failed to scan data")
	}

	result := sortingservice.SortAndCompareSlices(containersSize, ballColorCounts)

	if err = writer.Write(result); err != nil {
		log.Panic().Err(err).Msg("failed to write result")
	}
}
