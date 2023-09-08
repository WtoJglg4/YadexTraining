package set

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DistinctNumbers() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	strByte, _, err := in.ReadLine()
	if err != nil {
		fmt.Fprintln(out, err)
		return
	}

	strNumbers := strings.Split(string(strByte), " ")
	if len(strNumbers) == 1 && strNumbers[0] == "" {
		fmt.Fprintln(out, 0)
		return
	}

	unique := make(map[int]struct{}, 0)
	for _, v := range strNumbers {
		number, _ := strconv.Atoi(v)
		unique[number] = struct{}{}
		// if _, ok := unique[number]; !ok {
		// 	unique[number] = struct{}{}
		// }
	}
	fmt.Fprintln(out, len(unique))
}

func DistinctNumbersTest() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	numbers := strings.Split(input, " ")
	//fmt.Println(numbers)

	set := make(map[int]struct{})
	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		set[num] = struct{}{}
	}
	//fmt.Println(set)
	fmt.Println(len(set))
}

func DistinctNumbersSet() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	numbers := strings.Split(input, " ")
	//fmt.Println(numbers)

	set := NewSet(10)
	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		set.Add(num)
	}
	var cnt int
	for _, bucket := range set {
		cnt += len(bucket)
	}
	//fmt.Println(set)
	fmt.Println(cnt)
}
