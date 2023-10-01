package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

func Tourism() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n)

	points := make([][]int, n)
	for i := range points {
		point := []int{0, 0}
		fmt.Fscan(in, &point[0], &point[1])
		points[i] = point
	}
	pointsReverse := make([][]int, n)
	for i := 0; i < n; i++ {
		pointsReverse[i] = points[n-1-i]
	}

	uphillPrefix := getPrefixUphill(points)
	uphillPrefixReverse := getPrefixUphill(pointsReverse)

	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		var beg, end, count int
		fmt.Fscan(in, &beg, &end)
		if beg > end {

			count = uphillPrefixReverse[n-end] - uphillPrefixReverse[n-beg]
		} else {

			count = uphillPrefix[end-1] - uphillPrefix[beg-1]
		}
		fmt.Fprintln(out, count)
	}

}

func getPrefixUphill(points [][]int) []int {
	prefix := make([]int, len(points)+1)
	for i := 0; i < len(points)-1; i++ {
		if points[i+1][1] > points[i][1] {
			prefix[i+1] = prefix[i] + points[i+1][1] - points[i][1]
		} else {
			prefix[i+1] = prefix[i]
		}
	}
	return prefix
}
