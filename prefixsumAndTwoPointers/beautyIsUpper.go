package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

func BeautyIsUpper() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	trees := make([]int, n)
	for i := range trees {
		fmt.Fscan(in, &trees[i])
	}

	deficit := make(map[int]struct{}, k)
	proficit := make(map[int]int)
	for i := 1; i <= k; i++ {
		deficit[i] = struct{}{}
	}

	left, right := 0, 1
	minLeft, minRight, minLen := 0, 1, 250001
	proficit[trees[left]] = 1
	delete(deficit, trees[left])

	for right < n {
		rightEl := trees[right]
		if _, ok := deficit[rightEl]; ok {
			proficit[rightEl] = 0
			delete(deficit, rightEl)
		}
		proficit[rightEl]++

		if len(deficit) == 0 {
			leng := right - left
			if leng < minLen {
				minLen = leng
				minRight = right
				minLeft = left
			}

			for len(deficit) == 0 {
				leftEl := trees[left]
				proficit[leftEl]--
				left++
				if proficit[leftEl] != 0 {
					leng := right - left
					if leng < minLen {
						minLen = leng
						minRight = right
						minLeft = left
					}
				} else {
					deficit[leftEl] = struct{}{}
					break
				}
			}
		}
		right++
	}
	fmt.Fprintln(out, minLeft+1, minRight+1)
}
