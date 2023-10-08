package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

func Conditioners() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	rooms := make([]int, n)
	trash := make([]int, n)
	for i := range rooms {
		fmt.Fscan(in, &rooms[i])
	}
	rooms, trash = MergeSortTwo(rooms, trash)
	trash[0]++

	var m int
	fmt.Fscan(in, &m)
	powers := make([]int, m)
	prices := make([]int, m)
	for i := range powers {
		fmt.Fscan(in, &powers[i], &prices[i])
	}

	// fmt.Fprintf(out, "rooms: %v\tpowers: %v\tprices: %v\n", rooms, powers, prices)
	prices, powers = MergeSortTwo(prices, powers)
	// fmt.Fprintf(out, "Sort powers: %v\tSort prices: %v\n", powers, prices)

	roomIndex, priceIndex, sum := 0, 0, 0
	for priceIndex < m {
		price, power := prices[priceIndex], powers[priceIndex]

		count := 0
		for roomIndex < n && rooms[roomIndex] <= power {
			count++
			roomIndex++
		}
		sum += price * count

		if roomIndex == n {
			break
		}

		priceIndex++
	}
	fmt.Fprintln(out, sum)
}

func MergeSortTwo(list1, list2 []int) ([]int, []int) {

	n := len(list1)
	if n == 1 {
		return list1, list2
	}
	var FirstHalf, SecondHalf []int
	var FirstHalf2, SecondHalf2 []int
	if n%2 == 0 {
		FirstHalf = list1[:n/2]
		SecondHalf = list1[n/2:]
		FirstHalf2 = list2[:n/2]
		SecondHalf2 = list2[n/2:]

	} else {
		FirstHalf = list1[:n/2+1]
		SecondHalf = list1[n/2+1:]
		FirstHalf2 = list2[:n/2+1]
		SecondHalf2 = list2[n/2+1:]
	}

	SortedFH, SortedFH2 := MergeSortTwo(FirstHalf, FirstHalf2)
	SortedSH, SortedSH2 := MergeSortTwo(SecondHalf, SecondHalf2)

	SortedList, SortedList2 := MergeTwo(SortedFH, SortedSH, SortedFH2, SortedSH2)
	return SortedList, SortedList2
}
