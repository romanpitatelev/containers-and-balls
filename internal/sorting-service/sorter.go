package sortingservice

import "sort"

func SortAndCompareSlices(containers []int, balls []int) string {
	sort.Ints(containers)
	sort.Ints(balls)

	for i := range containers {
		if containers[i] != balls[i] {
			return "no"
		}
	}

	return "yes"
}
