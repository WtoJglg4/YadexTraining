package eventsort

import "fmt"

const (
	x1 = -1
	x2 = 1
)

func Points() {
	var n, m int
	fmt.Scan(&n, &m)
	xLine := make([][]int, 0)
	fmt.Println(n, m)
	for i := 0; i < n; i++ {
		var b, e int
		fmt.Scan(&b, &e)
		xLine = append(xLine, []int{b, x1})
		xLine = append(xLine, []int{e, x2})
	}
	fmt.Println(xLine)
}
