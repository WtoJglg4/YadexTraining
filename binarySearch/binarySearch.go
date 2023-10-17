package binsearch

import "fmt"

//to find first matching element
func LeftBinSearch(nums []int, l, r int, find int, check func(int, int) bool) int {
	for l < r {
		m := (l + r) / 2

		//if true => go to left part(including m)
		if check(nums[m], find) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

//to find last matching element
func RightBinSearch(nums []int, l, r int, find int, check func(int, int) bool) int {
	for l < r {
		m := (l + r + 1) / 2
		//if true => go to right part(including m)
		if check(nums[m], find) {
			l = m
		} else {
			r = m - 1
		}
		fmt.Printf("l: %d, r: %d, %v, %v\n", l, r, m, check(nums[m], find))
	}
	return r
}

func CheckLess(m, checkparams int) bool {
	return checkparams <= m
}

func CheckGreater(m, checkparams int) bool {
	return checkparams >= m
}
