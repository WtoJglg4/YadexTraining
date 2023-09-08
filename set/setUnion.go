package set

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UnionTwoSets() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	strNumbers := strings.Split(input, " ")
	empty1 := false
	empty2 := false
	if input == "" {
		empty1 = true
	}

	union := make(map[int]struct{})
	//fmt.Println(union, strNumbers)
	for _, v := range strNumbers {
		num, _ := strconv.Atoi(v)
		union[num] = struct{}{}
	}

	scanner.Scan()
	input = scanner.Text()
	strNumbers = strings.Split(input, " ")
	if input == "" {
		empty2 = true
	}
	if empty1 && empty2 {
		return
	}
	list := make([]int, len(strNumbers))
	for i, v := range strNumbers {
		list[i], _ = strconv.Atoi(v)
	}
	list = QuickSort(list)
	//fmt.Println(len(list), len(union), list, union)
	if len(list) == 0 && len(union) == 0 {
		return
	}
	for _, v := range list {
		if _, ok := union[v]; ok {
			fmt.Printf("%d ", v)
		}
	}
}

func UnionTwoSetsFORTEST(strNumbers, n2 []string) {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()

	//strNumbers := strings.Split(input, " ")
	unique := make(map[int]bool)
	for _, v := range strNumbers {
		num, _ := strconv.Atoi(v)
		unique[num] = false
	}

	// scanner.Scan()
	// input = scanner.Text()
	// strNumbers = strings.Split(input, " ")
	strNumbers = n2
	for _, v := range strNumbers {
		num, _ := strconv.Atoi(v)
		if _, ok := unique[num]; ok {
			unique[num] = true
		}
	}
	numbers := make([]int, 0)
	for k, v := range unique {
		if v {
			numbers = append(numbers, k)
		}
	}
	if len(numbers) == 0 {
		fmt.Printf("")
		return
	}
	numbers = QuickSort(numbers)
	for _, v := range numbers {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n ")
}

func QuickSort(list []int) []int {
	n := len(list)
	if n <= 1 {
		return list
	}

	var c_small, c_large, c_sorted []int
	for i, v := range list {
		if v < list[0] {
			c_small = append(c_small, list[i])
		}
		if v > list[0] {
			c_large = append(c_large, list[i])
		}
	}
	if len(c_small) != 0 {
		c_sorted = append(c_sorted, QuickSort(c_small)...)
	}
	c_sorted = append(c_sorted, list[0])
	if len(c_large) != 0 {
		c_sorted = append(c_sorted, QuickSort(c_large)...)
	}

	return c_sorted
}
