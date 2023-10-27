package binsearch

import "fmt"

func mergeL(s1, s2 []int) int {
	l, left, right := len(s1), 0, 0
	merge := make([]int, l)
	for i := 0; i < l; i++ {
		if s1[left] < s2[right] {
			merge[i] = s1[left]
			left++
		} else {
			merge[i] = s2[right]
			right++
		}
	}
	// fmt.Println("merge: ", merge)
	return merge[l-1]
}

func Mediana() {
	var n, l int
	fmt.Scan(&n, &l)
	numsSlice := make([][]int, n)
	for i := range numsSlice {
		numsSlice[i] = make([]int, l)
		for j := range numsSlice[i] {
			fmt.Scan(&numsSlice[i][j])
		}
	}
	// fmt.Println(numsSlice)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fmt.Println(mergeL(numsSlice[i], numsSlice[j]))
		}
	}

}
