package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"os"
)

type Letter struct {
	count   int
	indexes []int
}

func Substring() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)

	maxLength, maxFirstIndex := 1, 0
	length, left, right := 1, 0, 1
	letters := make(map[rune]Letter)
	var str string
	fmt.Fscan(in, &str)
	strRune := []rune(str)
	letters[strRune[left]] = Letter{1, []int{0}}

	for right < n {
		rightEl := strRune[right]
		if _, ok := letters[rightEl]; !ok {
			letters[rightEl] = Letter{0, []int{}}
		}

		if v := letters[rightEl]; v.count < k {
			letters[rightEl] = Letter{v.count + 1, append(v.indexes, right)}

			length = right - left + 1
			if length > maxLength {
				maxLength = length
				maxFirstIndex = left
			}
			right++
		} else {
			nexLeft := v.indexes[0] + 1
			for left < nexLeft {
				l := letters[strRune[left]]
				letters[strRune[left]] = Letter{l.count - 1, l.indexes[1:]}
				left++
			}

			letters[rightEl] = Letter{v.count, append(v.indexes[1:], right)}

			right++
		}
	}
	fmt.Fprintln(out, maxLength, maxFirstIndex+1)
}

func SubstringSlow(n, k int, str string) (int, int) {
	maxLen, maxFirstInd := 0, 0
	strRune := []rune(str)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			letters := make(map[rune]int)
			length := 0
			for k := i; k <= j; k++ {
				letters[strRune[k]]++
				length++
			}
			flag := true
			for _, v := range letters {
				if v > k {
					flag = false
					break
				}
			}
			if flag && length > maxLen {
				maxLen = length
				maxFirstInd = i
			}
		}
	}
	return maxLen, maxFirstInd + 1
}
