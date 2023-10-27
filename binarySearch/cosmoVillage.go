package binsearch

import (
	"bufio"
	"fmt"
	"github/YandexTraining/mathem"
	"os"
)

func rightCosmoSearch(n, a, b, w, h int) int {
	w, h = mathem.Min(w, h), mathem.Max(w, h)
	a, b = mathem.Min(a, b), mathem.Max(a, b)
	l, r := 1, mathem.Max((h-n*b)/(2*n), (w-n*a)/(2*n))
	for l < r {
		m := (l + r + 1) / 2
		realW, realH := 2*m+a, 2*m+b
		if (h/realH)*(w/realW) <= n {
			l = m
		} else {
			r = m - 1
		}
	}
	return r
}

func CosmoVillage() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a, b, w, h int
	fmt.Fscan(in, &n, &a, &b, &w, &h)
	fmt.Fprintln(out, rightCosmoSearch(n, a, b, w, h))
}
