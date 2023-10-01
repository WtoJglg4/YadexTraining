package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

func NumberSum() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k, n int
	fmt.Fscan(in, &n, &k)

	numbers := make([]int, n)
	for i := range numbers {
		fmt.Fscan(in, &numbers[i])
	}
	prefix := getPrefix(numbers)

	left, right, count := 0, 1, 0
	for right < len(prefix) {
		diff := prefix[right] - prefix[left]
		if diff < k {
			right++
		} else if diff == k {
			count++
			right++
			left++
		} else {
			left++
		}
	}
	fmt.Fprintln(out, count)

}

func getPrefix(list []int) []int {
	prefix := make([]int, len(list)+1)
	for i := range list {
		prefix[i+1] = list[i] + prefix[i]
	}
	return prefix
}
