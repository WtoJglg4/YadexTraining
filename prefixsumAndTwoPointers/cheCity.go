package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

func CheCity() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, r int
	fmt.Fscan(in, &n, &r)
	monuments := make([]int, n)
	for i := range monuments {
		fmt.Fscan(in, &monuments[i])
	}

	left, right, count := 0, 1, 0
	for right < n {
		diff := monuments[right] - monuments[left]
		if diff > r {
			// fmt.Fprintln(out, left, right)
			count += n - right
			left++
		} else {
			right++
		}
	}
	fmt.Fprint(out, count)
}
