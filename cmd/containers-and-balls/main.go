package main

import (
	"sort"

	scannerservice "github.com/romanpitatelev/containers-and-balls/internal/scanner-service"
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

	result := compareSlices(containersSize, ballColorCounts)

	if err = writer.Write(result); err != nil {
		log.Panic().Err(err).Msg("failed to write result")
	}
}

func compareSlices(containers []int, balls []int) string {
	sort.Ints(containers)
	sort.Ints(balls)

	for i := range containers {
		if containers[i] != balls[i] {
			return "no"
		}
	}

	return "yes"
}
