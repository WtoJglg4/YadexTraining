package prefixsumAndTwoPointers

func MakePrefixSum(nums []int) []int {
	prefixsum := make([]int, len(nums)+1)
	prefixsum[0] = 0
	for i := 1; i < len(prefixsum); i++ {
		prefixsum[i] = prefixsum[i-1] + nums[i-1]
	}
	return prefixsum
}

// range some query
func RSQ(prefixsum []int, l, r int) int {
	return prefixsum[r] - prefixsum[l]
}
