package testing

import (
	"fmt"
	"math/rand"
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
	nums := make([]int, 20)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		for i := 0; i < len(nums); i++ {
			nums[i] = r1.Intn(201) - 100
		}
		fmt.Println(nums)
	}

}
