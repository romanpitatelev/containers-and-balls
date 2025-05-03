package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minN             = 1
	maxN             = 100
	minNumberOfBalls = 0
	maxNumberOfBalls = 1000000000
)

var (
	ErrInvalidN             = errors.New("n is out of acceptable range")
	ErrWrongRowLength       = errors.New("invalid row length")
	ErrInvalidNumberOfBalls = errors.New("number of balls is out of acceptable range")
)

func ReadNumber() (int, error) {
	input := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscanln(input, &n); err != nil {
		return 0, fmt.Errorf("scanning error: %w", err)
	}

	if err := validateNumber(n); err != nil {
		return 0, fmt.Errorf("n validation error: %w", err)
	}

	return n, nil
}

func ReadContainersBals(number int) ([]int, []int, error) {
	input := bufio.NewReader(os.Stdin)

	containersSize := make([]int, number)
	ballColorCounts := make([]int, number)

	for i := range number {
		line, err := input.ReadString('\n')
		if err != nil {
			return nil, nil, fmt.Errorf("error reading line: %w", err)
		}

		line = strings.TrimSpace(line)
		strContainerSlice := strings.Fields(line)

		if len(strContainerSlice) != number {
			return nil, nil, fmt.Errorf("error: %w; line %d has %d elements, expected %d", ErrWrongRowLength, i+1, len(strContainerSlice), number)
		}

		var ballsInContainer int

		for idx, ballsStr := range strContainerSlice {
			balls, err := strconv.Atoi(ballsStr)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing integer at container %d: %w", idx+1, err)
			}

			if err := validateContainersBalls(balls); err != nil {
				return nil, nil, fmt.Errorf("balls validation error: %w", ErrInvalidNumberOfBalls)
			}

			ballsInContainer += balls

			ballColorCounts[idx] += balls
		}

		containersSize[i] = ballsInContainer
	}

	return containersSize, ballColorCounts, nil
}

func validateNumber(n int) error {
	if n < minN || n > maxN {
		return ErrInvalidN
	}

	return nil
}

func validateContainersBalls(number int) error {
	if number < minNumberOfBalls || number > maxNumberOfBalls {
		return ErrInvalidNumberOfBalls
	}

	return nil
}
