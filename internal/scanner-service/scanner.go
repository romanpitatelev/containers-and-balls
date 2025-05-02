package scannerservice

import (
	"bufio"
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

type Config struct {
	nMin             int
	nMax             int
	numberOfBallsMin int
	numberOfBallsMax int
}

type Scanner struct {
	cfg Config
}

func New() *Scanner {
	return &Scanner{
		cfg: Config{
			nMin:             minN,
			nMax:             maxN,
			numberOfBallsMin: minNumberOfBalls,
			numberOfBallsMax: maxNumberOfBalls,
		},
	}
}

func (s *Scanner) Scan() ([]int, []int, error) {
	input := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscanln(input, &n); err != nil {
		return nil, nil, fmt.Errorf("scanning error: %w", err)
	}

	if err := s.validateN(n); err != nil {
		return nil, nil, fmt.Errorf("n validation error: %w", err)
	}

	containersSize := make([]int, n)
	ballColorCounts := make([]int, n)

	for i := range n {
		line, err := input.ReadString('\n')
		if err != nil {
			return nil, nil, fmt.Errorf("error reading line: %w", err)
		}

		line = strings.TrimSpace(line)
		strContainerSlice := strings.Fields(line)

		if len(strContainerSlice) != n {
			return nil, nil, fmt.Errorf("error: %w; line %d has %d elements, expected %d", ErrWrongRowLength, i+1, len(strContainerSlice), n)
		}

		var ballsInContainer int

		for idx, ballsStr := range strContainerSlice {
			balls, err := strconv.Atoi(ballsStr)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing integer at container %d: %w", idx+1, err)
			}

			if err := s.validateNumberOfBalls(balls); err != nil {
				return nil, nil, fmt.Errorf("balls validation error: %w", ErrInvalidNumberOfBalls)
			}

			ballsInContainer += balls

			ballColorCounts[idx] += balls
		}

		containersSize[i] = ballsInContainer
	}

	return containersSize, ballColorCounts, nil
}

func (s *Scanner) validateN(n int) error {
	if n < s.cfg.nMin || n > s.cfg.nMax {
		return ErrInvalidN
	}

	return nil
}

func (s *Scanner) validateNumberOfBalls(number int) error {
	if number < s.cfg.numberOfBallsMin || number > s.cfg.numberOfBallsMax {
		return ErrInvalidNumberOfBalls
	}

	return nil
}
