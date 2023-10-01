package prefixsumAndTwoPointers

import "fmt"

func Soccer() {
	players := []int{1, 1, 3, 3, 4, 6, 11}
	bestsum, cursum, last := 0, 0, 0
	for first := range players {
		for last < len(players) && (last == first || players[first]+players[first+1] >= players[last]) {
			cursum += players[last]
			last++
		}
		bestsum = max(bestsum, cursum)
		cursum -= players[first]
	}
	fmt.Print(bestsum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
