package t

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProductThree() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strNums := strings.Split(input, " ")
	nums := make([]int, len(strNums))
	for i, v := range strNums {
		nums[i], _ = strconv.Atoi(v)
	}
	fmt.Println(nums)
}
