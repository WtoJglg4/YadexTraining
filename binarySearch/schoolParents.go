package binsearch

func SchoolParents(nums []int, l, n, k int) int {
	r := n - 1
	for l < r {
		m := (l + r) / 2
		if 2*nums[m] >= n-3*k {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func YandexInterview(n, k int) int {
	tasks := make([]int, n)
	tasks[0] = k
	for i := 1; i < n; i++ {
		tasks[i] = tasks[i-1] + 1
	}
	return LeftBinSearch(tasks, 0, n-1, n, CheckLess)
}
