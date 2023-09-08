package testing

import (
	"math/rand"
	"strconv"
	"time"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Testing() {
	nums := make([]string, 10)
	nums1 := make([]string, 10)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	s := 0
	for s < 2 {
		for i := 0; i < len(nums); i++ {
			nums[i] = strconv.Itoa(r1.Intn(201) - 100)

		}
		for i := 0; i < len(nums); i++ {
			nums1[i] = strconv.Itoa(r1.Intn(201) - 100)
		}
		///fmt.Println(nums)
		//fmt.Println(nums1)
		//set.UnionTwoSets(nums, nums1)
		s++
	}

}
