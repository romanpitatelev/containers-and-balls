package sorter

import "sort"

func SortAndCompareSlices(containers []int, balls []int) bool {
	sort.Ints(containers)
	sort.Ints(balls)

	for i := range containers {
		if containers[i] != balls[i] {
			return false
		}
	}

	return true
}
