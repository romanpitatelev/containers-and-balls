package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	nMin             = 1
	nMax             = 100
	numberOfBallsMin = 0
	numberOfBallsMax = 1000000000
)

func main() {
	var (
		input  *bufio.Reader
		output *bufio.Writer
	)

	input = bufio.NewReader(os.Stdin)
	output = bufio.NewWriter(os.Stdout)
	defer output.Flush()

	var n int

	if _, err := fmt.Fscanln(input, &n); err != nil {
		fmt.Fprintln(output, "Error reading n:", err)

		return
	}

	if err := validateN(n); err != nil {
		fmt.Fprintln(output, "n validation error:", err)

		return
	}

	containersSize := make([]int, n)
	ballColorCounts := make([]int, n)

	for i := 0; i < n; i++ {
		line, err := input.ReadString('\n')
		if err != nil {
			fmt.Fprintln(output, "Error reading line:", err)

			return
		}

		line = strings.TrimSpace(line)
		strContainerSlice := strings.Fields(line)
		if len(strContainerSlice) != n {
			fmt.Fprintf(output, "Error: line %d has %d elements, expected %d", i+1, len(strContainerSlice), n)

			return
		}

		var ballsInContainer int

		for idx, ballsStr := range strContainerSlice {
			balls, err := strconv.Atoi(ballsStr)
			if err != nil {
				fmt.Fprintf(output, "Error parsing integer at container %d: %v", idx+1, err)

				return
			}

			if err := validateNumberBalls(balls); err != nil {
				fmt.Fprintln(output, "number of balls validation error:", err)

				return
			}

			ballsInContainer += balls

			ballColorCounts[idx] += balls
		}

		containersSize[i] = ballsInContainer

	}

	result := compareSlices(containersSize, ballColorCounts)

	fmt.Fprintln(output, result)
}

func compareSlices(containers []int, balls []int) string {
	sort.Ints(containers)
	sort.Ints(balls)

	for i := 0; i < len(containers); i++ {
		if containers[i] != balls[i] {
			return "no"
		}
	}

	return "yes"
}

func validateN(n int) error {
	if n < nMin || n > nMax {
		return fmt.Errorf("n is out of acceptable range")
	}

	return nil
}

func validateNumberBalls(number int) error {
	if number < numberOfBallsMin || number > numberOfBallsMax {
		return fmt.Errorf("number of balls is out of acceptable range")
	}

	return nil
}
