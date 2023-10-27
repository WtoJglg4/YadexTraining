package binsearch

import "fmt"

func cutLines(lines []int, m int) int {
	sum := 0
	for _, v := range lines {
		sum += v / m
	}
	return sum
}

func linesRBS(n, k int, lines []int) int {
	l, r := 0, 10000000
	for l < r {
		m := (l + r + 1) / 2
		fmt.Println(l, r, m, cutLines(lines, m))
		if cutLines(lines, m) >= k {
			l = m
		} else {
			r = m - 1
		}
	}
	return r
}

func Lines() {
	var n, k int
	fmt.Scan(&n, &k)
	lines := make([]int, n)
	for i := range lines {
		fmt.Scan(&lines[i])
	}
	fmt.Println(linesRBS(n, k, lines))
}
