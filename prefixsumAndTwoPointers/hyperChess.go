package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

func HyperChess() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(in, &nums[i])
	}

	nums = MergeSort(nums)

	count, left, right := 0, 0, 1
	numsCount := make(map[int]int)
	numsCount[nums[0]]++
	for right < n {
		fmt.Fprintln(out, left, right)
		numsCount[nums[right]]++
		if nums[right]/nums[left] <= k {

			// fmt.Fprintln(out, left, right, right-left, Fact(right-left), count)
			right++
		} else {
			fmt.Fprintln(out, "hui", left, right)
			numsCount[nums[right]]--

			for nums[right]/nums[left] > k {
				fmt.Fprintln(out, "penis", left, right)
				numsCount[nums[left]]--
				left++
			}
			right--
			count += Fact(right-left) / IntervalCount(numsCount)
			// left++
			right++
		}

	}
	fmt.Fprintln(out, count)
}

func MergeSort(list []int) []int {

	n := len(list)
	if n == 1 {
		return list
	}
	var FirstHalf, SecondHalf []int
	if n%2 == 0 {
		FirstHalf = list[:n/2]
		SecondHalf = list[n/2:]

	} else {
		FirstHalf = list[:n/2+1]
		SecondHalf = list[n/2+1:]
	}

	SortedFH := MergeSort(FirstHalf)
	SortedSH := MergeSort(SecondHalf)

	SortedList := Merge(SortedFH, SortedSH)
	return SortedList
}

func Fact(num int) int {
	if num < 2 {
		return 1
	}
	return num * Fact(num-1)
}

func IntervalCount(numsCount map[int]int) int {
	count := 0
	for _, v := range numsCount {
		count += Fact(v)
	}
	return count
}
