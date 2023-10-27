package binsearch

import "fmt"

func clumbaRBS(n, m, t int64) int64 {
	var l, r, a, b int64
	l, r = 0, t
	for l < r {
		middle := (l + r + 1) / 2
		a, b = (n - 2*middle), (m - 2*middle)
		s := a * b
		fmt.Println(l, r, middle, a, b, s)
		if n*m-s <= t && a > 0 && b > 0 {
			l = middle
		} else {
			r = middle - 1
		}
	}
	return r
}

func SqareSS() {
	var n, m, t int64
	fmt.Scan(&n, &m, &t)
	fmt.Println(clumbaRBS(n, m, t))
}

func Max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
