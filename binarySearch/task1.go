package binsearch

import (
	"bufio"
	"fmt"
	"github/YandexTraining/mathem"
	"os"
)

func lbs(num int, nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < num {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func rbs(num int, nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r + 1) / 2
		if nums[m] > num {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func Task1() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, num int
	fmt.Fscan(in, &n, &k)
	nums1 := make([]int, n)
	for i := range nums1 {
		fmt.Fscan(in, &nums1[i])
	}
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &num)
		ind1 := lbs(num, nums1)
		ind2 := rbs(num, nums1)
		if mathem.Abs(nums1[ind1]-num) < mathem.Abs(nums1[ind2]-num) {
			fmt.Fprintln(out, nums1[ind1])
		} else {
			fmt.Fprintln(out, nums1[ind2])
		}
	}
}
