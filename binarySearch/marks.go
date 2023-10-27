package binsearch

import "fmt"

func avgMarks(a, b, c, d int) bool {
	return 2*(2*a+3*b+4*c+5*d) >= 7*(a+b+c+d)
}

func markLBS(a, b, c, maxFive int) int {
	l, r := 0, maxFive
	for l < r {
		m := (l + r) / 2
		// fmt.Println(l, r, m, avgMarks(a, b, c, m))
		if avgMarks(a, b, c, m) {
			r = m
		} else {
			l = m + 1
		}

	}
	return l
}

func Marks() {
	var two, three, four int
	fmt.Scan(&two, &three, &four)
	maxFive := 5*two + 3*three + four
	fmt.Println(markLBS(two, three, four, maxFive))
}
