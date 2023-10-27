package binsearch

import (
	"fmt"
	"github/YandexTraining/mathem"
)

func easylbs(N, x, y int) int {
	l, r := 1, N*mathem.Max(x, y)
	for l < r {
		m := (l + r) / 2
		if m/x+m/y >= N {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func EasyTask() {
	var N, x, y int
	fmt.Scan(&N, &x, &y)
	N--
	total := mathem.Min(x, y)
	if N < 1 {
		fmt.Println(total)
		return
	}
	fmt.Println(total + easylbs(N, x, y))
}
