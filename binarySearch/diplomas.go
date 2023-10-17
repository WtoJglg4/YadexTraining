package binsearch

import (
	"bufio"
	"fmt"
	"github/YandexTraining/mathem"
	"os"
)

func sqaresearch(w, h, n int) int {
	l, r := 0, mathem.Max(w, h)*n
	for l < r {
		m := (l + r) / 2
		if (m/w)*(m/h) >= n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func Diplomas() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var w, h, n int
	fmt.Fscan(in, &w, &h, &n)
	fmt.Fprintln(out, sqaresearch(w, h, n))
}
