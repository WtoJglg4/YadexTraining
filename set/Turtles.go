package set

import "fmt"

func Turtles() {
	var n, a, b, count int
	fmt.Scan(&n)
	seq := make([]bool, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		if a+b != n-1 || b < 0 || b >= n {
			continue
		}
		if !seq[b] {
			seq[b] = true
			count++
		}

	}
	fmt.Println(count)
}
