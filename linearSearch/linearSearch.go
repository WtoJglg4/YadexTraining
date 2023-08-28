package linear

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sapper() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	field := make([][]int, n)
	for i := range field {
		field[i] = make([]int, m)
	}

	for l := 0; l < k; l++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		field[x-1][y-1] = -1
	}

	printField(field)

}

func printField(field [][]int) {
	for i, v := range field {
		for j, w := range v {
			if w == -1 {
				fmt.Printf("%v ", "*")
				continue
			}
			var mineCnt int

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}

					if j == 0 {
						if y < 0 || i == 0 && x < 0 || i == len(field)-1 && x > 0 {
							continue
						}
					}
					if j == len(field[i])-1 {
						if y > 0 || i == 0 && x < 0 || i == len(field)-1 && x > 0 {
							continue
						}
					}
					if i == 0 {
						if x < 0 || j == 0 && y < 0 || j == len(field[i])-1 && y > 0 {
							continue
						}
					}
					if i == len(field)-1 {
						if x > 0 || j == 0 && y < 0 || j == len(field[i])-1 && y > 0 {
							continue
						}
					}

					if field[i+x][j+y] == -1 {
						mineCnt++
					}
				}
			}
			fmt.Printf("%d ", mineCnt)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func maxProduct2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strNum := scanner.Text()

	numsString := strings.Split(strNum, " ")
	//fmt.Println(numsString)
	nums := make([]int, len(numsString))
	for i, v := range numsString {
		nums[i], _ = strconv.Atoi(v)
	}

	var maxPos, minNeg, index1, index2 int
	maxPos = -1
	minNeg = 1
	for i, v := range nums {
		if v < minNeg && v <= 0 {
			minNeg = v
			index2 = i
		}
		if v > maxPos && v >= 0 {
			maxPos = v
			index1 = i
		}
		// fmt.Println(maxPos, minNeg)
		// fmt.Println(index1, index2)
	}
	// fmt.Printf("\n")
	var pos, neg int
	//index3, index4 := -1, -1
	pos, neg = -1, 1
	for i, v := range nums {
		if v < neg && v <= 0 && i != index2 {
			neg = v
			//index4 = i
		}
		if v > pos && v >= 0 && i != index1 {
			pos = v
			//index3 = i
		}
		// fmt.Println(pos, neg)
		// fmt.Println(index3, index4)
	}

	//fmt.Printf("maxPos:%d minNeg:%d pos:%d neg:%d\n", maxPos, minNeg, pos, neg)
	if pos == -1 {
		if neg == 1 {
			fmt.Println(math.Min(float64(maxPos), float64(minNeg)), math.Max(float64(maxPos), float64(minNeg)))
			return
		} else {
			fmt.Println(math.Min(float64(minNeg), float64(neg)), math.Max(float64(minNeg), float64(neg)))
			return
		}
	}
	if maxPos*pos > minNeg*neg || neg == 1 {
		fmt.Println(math.Min(float64(maxPos), float64(pos)), math.Max(float64(maxPos), float64(pos)))
		// return min(maxPos, pos), max(maxPos, pos)
	} else {
		fmt.Println(math.Min(float64(minNeg), float64(neg)), math.Max(float64(minNeg), float64(neg)))
		// return min(minNeg, neg), max(minNeg, neg)
	}
}

// func maxProduct1(nums []int) (int, int) {
// func maxProduct1() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	strNum := scanner.Text()

// 	numsString := strings.Split(strNum, " ")
// 	//fmt.Println(numsString)
// 	nums := make([]int, len(numsString))
// 	for i, v := range numsString {
// 		nums[i], _ = strconv.Atoi(v)
// 	}

// 	var maxPos, minNeg, index1, index2 int
// 	maxPos = -1
// 	minNeg = 1
// 	for i, v := range nums {
// 		if v < minNeg && v <= 0 {
// 			minNeg = v
// 			index2 = i
// 		}
// 		if v > maxPos && v >= 0 {
// 			maxPos = v
// 			index1 = i
// 		}
// 		// fmt.Println(maxPos, minNeg)
// 		// fmt.Println(index1, index2)
// 	}
// 	// fmt.Printf("\n")
// 	var pos, neg int
// 	//index3, index4 := -1, -1
// 	pos, neg = -1, 1
// 	for i, v := range nums {
// 		if v < neg && v <= 0 && i != index2 {
// 			neg = v
// 			//index4 = i
// 		}
// 		if v > pos && v >= 0 && i != index1 {
// 			pos = v
// 			//index3 = i
// 		}
// 		// fmt.Println(pos, neg)
// 		// fmt.Println(index3, index4)
// 	}

// 	//fmt.Printf("maxPos:%d minNeg:%d pos:%d neg:%d\n", maxPos, minNeg, pos, neg)
// 	if pos == -1 {
// 		if neg == 1 {
// 			testing.min()
// 			fmt.Println(testing.min(maxPos, minNeg), testing.max(maxPos, minNeg))
// 			return
// 		} else {
// 			fmt.Println(testing.min(minNeg, neg), testing.max(minNeg, neg))
// 			return
// 		}
// 	}
// 	if maxPos*pos > minNeg*neg || neg == 1 {
// 		fmt.Println(min(maxPos, pos), max(maxPos, pos))
// 		// return min(maxPos, pos), max(maxPos, pos)
// 	} else {
// 		fmt.Println(min(minNeg, neg), max(minNeg, neg))
// 		// return min(minNeg, neg), max(minNeg, neg)
// 	}

// }

// func maxProductSLOW(nums []int) (int, int) {
// 	maxRes := nums[0] * nums[1]
// 	m1, m2 := min(nums[0], nums[1]), max(nums[0], nums[1])
// 	for i, v := range nums {
// 		for j, w := range nums {
// 			if v*w > maxRes && i != j {
// 				m1, m2 = min(v, w), max(v, w)
// 				maxRes = v * w
// 			}
// 		}
// 	}
// 	fmt.Println(m1, m2)
// 	return m1, m2
// }

// func maxProduct(nums []int) (int, int) {
// 	// scanner := bufio.NewScanner(os.Stdin)
// 	// scanner.Scan()
// 	// strs := strings.Split(scanner.Text(), " ")
// 	// nums := make([]int, len(strs))
// 	// for i := 0; i < len(nums); i++ {
// 	// 	nums[i], _ = strconv.Atoi(strs[i])
// 	// }

// 	var max1, max2, index int
// 	max1 = nums[0]
// 	index = 0
// 	max2 = math.MinInt
// 	for i, v := range nums {
// 		if v > max1 {
// 			max1 = v
// 			index = i
// 		}
// 	}
// 	for i, v := range nums {
// 		if v > max2 && i != index {
// 			max2 = v
// 		}
// 	}

// 	var min1, min2 int
// 	min1 = nums[0]
// 	index = 0
// 	min2 = math.MaxInt
// 	for i, v := range nums {
// 		if v < min1 {
// 			min1 = v
// 			index = i
// 		}
// 	}
// 	for i, v := range nums {
// 		if v < min2 && i != index {
// 			min2 = v
// 		}
// 	}

// 	if max1*max2 > min1*min2 {
// 		// fmt.Println(min(max2, max1), max(max2, max1))
// 		return min(max2, max1), max(max2, max1)
// 	} else {
// 		// fmt.Println(min(min1, min2), max(min1, min2))
// 		return min(min1, min2), max(min1, min2)
// 	}

// }

func symmetrySequence() {
	var n int
	fmt.Scan(&n)
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
	}

	sym := true
	dif := 0
	eq := len(seq) / 2
	for i := 0; i < eq; i++ {
		if seq[i] != seq[len(seq)-1-i] {
			sym = false
			break
		}
	}
	if sym {
		fmt.Println(0)
		return
	}

	for i := eq; i < len(seq); i++ {
		sym = true
		for j := 1; j < eq+1; j++ {
			//fmt.Println(i, j)
			if i+j >= len(seq) {
				dif = j
				break
			}
			if seq[i+j] != seq[i-j] {
				sym = false
				break
			}
		}
		//fmt.Println(sym, "for1")
		if sym {
			fmt.Println(i - dif + 1)
			for j := i - dif; j > -1; j-- {
				fmt.Print(seq[j], " ")
			}
			return
		}

		sym = true
		for j := 0; j < eq+1; j++ {
			//fmt.Println(i, j)
			if i+j+1 >= len(seq) {
				dif = j
				break
			}
			if seq[i+j+1] != seq[i-j] {
				sym = false
				break
			}
		}
		//fmt.Println(sym, "for2", dif)
		if sym {
			fmt.Println(i - dif + 1)
			for j := i - dif; j > -1; j-- {
				fmt.Print(seq[j], " ")
			}
			return
		}

	}
}

func cowsPattyChamp() {
	var n int
	fmt.Scan(&n)
	results := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&results[i])
	}

	maxRes := results[0]
	maxIndex := 0
	for i := 1; i < n; i++ {
		if results[i] > maxRes {
			maxRes = results[i]
			maxIndex = i
		}
	}

	maxRes = 0
	for i := maxIndex + 1; i < n-1; i++ {
		if results[i]%10 == 5 && results[i+1] < results[i] {
			if i == maxIndex+1 {
				maxRes = results[i]
			}
			if results[i] > maxRes {
				maxRes = results[i]
			}
		}
	}
	if maxRes == 0 {
		fmt.Println(maxRes)
		return
	}

	maxIndex = 0
	for i := 0; i < n; i++ {
		if results[i] > maxRes {
			maxIndex++
		}
	}
	fmt.Println(maxIndex + 1)
}

func neighbours() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), " ")
	nums := make([]int, len(strs))
	for i := 0; i < len(nums); i++ {
		nums[i], _ = strconv.Atoi(strs[i])
	}

	var cnt int
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}

// func closeNumList() {
// 	var n, m int
// 	var minDiff float64
// 	fmt.Scan(&n)
// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&nums[i])
// 	}
// 	fmt.Scan(&m)
// 	for i, v := range nums {
// 		if m == v {
// 			fmt.Println(m)
// 			return
// 		}
// 		if i == 0 {
// 			minDiff = math.Abs(float64(m - v))
// 			n = v
// 		} else {
// 			diff := math.Abs(float64(m - v))
// 			if diff < minDiff {
// 				minDiff = diff
// 				n = v
// 			}
// 		}

// 	}
// 	fmt.Println(n)
// }

// const (
// 	cons  = "CONSTANT"
// 	asc   = "ASCENDING"
// 	wasc  = "WEAKLY ASCENDING"
// 	desc  = "DESCENDING"
// 	wdesc = "WEAKLY DESCENDING"
// 	rand  = "RANDOM"
// 	stop  = -2000000000
// )

// func sequence() {
// 	state := ""
// 	var prev, cur int64
// 	fmt.Scan(&prev)
// 	if prev == stop {
// 		fmt.Println(state)
// 		return
// 	}
// 	for {
// 		fmt.Scan(&cur)
// 		if cur == stop {
// 			break
// 		}
// 		if cur > prev {
// 			if state == cons {
// 				state = wasc
// 			} else if state == desc || state == wdesc {
// 				state = rand
// 				break
// 			} else if state == "" {
// 				state = asc
// 			}
// 		} else if cur < prev {
// 			if state == cons {
// 				state = wdesc
// 			} else if state == asc || state == wasc {
// 				state = rand
// 				break
// 			} else if state == "" {
// 				state = desc
// 			}
// 		} else {
// 			if state == asc {
// 				state = wasc
// 			} else if state == desc {
// 				state = wdesc
// 			} else if state == "" {
// 				state = cons
// 			}
// 		}
// 		prev = cur
// 	}
// 	fmt.Println(state)
// }

func numsList() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	// fmt.Println(str)
	strs := strings.Split(str, " ")

	nums := make([]int, len(strs))
	for i, v := range strs {
		nums[i], _ = strconv.Atoi(v)
	}

	flag := true
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func alloy1() {
	var N, K, M, count int
	fmt.Scan(&N, &K, &M)

	for N >= K {
		count += N / K
		N = N/K*M + N%K
	}

	fmt.Println(count)
}

func alloy() {
	var n, k, m, mk, cnt int
	fmt.Scan(&n, &k, &m)
	k_div_m := k / m
	k_mod_m := k % m
	for n/k > 0 {
		//fmt.Println(n, k, m, mk)
		mk = n / k
		cnt += mk * k_div_m
		n = n%k + mk*k_mod_m
		//fmt.Println(n, k, m, mk, cnt)
	}
	fmt.Println(cnt)
}

func systemLinearEq() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)
	if a == 0 {
		if b == 0 {
			if e == 0 {
				if c == 0 {
					if d == 0 {
						if f == 0 {
							fmt.Println(5)
							return
						}
						fmt.Println(0)
						return
					}
					fmt.Println(4, f/d)
					return
				}
				if d == 0 {
					fmt.Println(3, f/c)
					return
				}
				fmt.Println(1, f/d, (-1)*c/d)
				return
			}
			fmt.Println(0)
			return
		}
		if c == 0 {
			if d == 0 {
				if f == 0 {
					fmt.Println(5)
					return
				}
				fmt.Println(0)
				return
			}
			if f/d == e/b {
				fmt.Println(4, e/b)
				return
			}
			fmt.Println(0)
			return
		}
		if d == 0 {
			fmt.Println(2, f/c, e/b)
			return
		}
		fmt.Println(3, f/c-(d*e)/(b*c))
		return
	}

	if c == 0 {
		if d == 0 {
			if f == 0 {
				if a == 0 {
					if b == 0 {
						if e == 0 {
							fmt.Println(5)
							return
						}
						fmt.Println(0)
						return
					}
					fmt.Println(4, e/b)
					return
				}
				if b == 0 {
					fmt.Println(3, e/a)
					return
				}
				fmt.Println(1, e/b, (-1)*a/b)
				return
			}
			fmt.Println(0)
			return
		}
		if a == 0 {
			if b == 0 {
				if e == 0 {
					fmt.Println(5)
					return
				}
				fmt.Println(0)
				return
			}
			if e/b == f/d {
				fmt.Println(4, f/d)
				return
			}
			fmt.Println(0)
			return
		}
		if b == 0 {
			fmt.Println(2, e/a, f/d)
			return
		}
		fmt.Println(3, e/a-(b*f)/(d*a))
		return
	}

	if a*d-c*b == 0 {
		if f*a == c*e {
			fmt.Println(1, (-1)*a/b, e/b)
			return
		}
		fmt.Println(0)
		return
	}
	y := (f*a - c*e) / (a*d - c*b)
	x := (e - b*y) / a
	fmt.Println(2, x, y)
}

// func uznikIF() {
// 	var a, b, c, d, e int
// 	fmt.Scan(&a, &b, &c, &d, &e)
// 	if c > a {
// 		a, c = c, a
// 	}
// 	if c > b {
// 		b, c = c, b
// 	}
// 	if b > a {
// 		b, a = a, b
// 	}

// 	if b > max(d, e) || c > min(d, e) {
// 		fmt.Println("NO")
// 		return
// 	}
// 	fmt.Println("YES")

// 	fmt.Println(a, b, c)
// 	if max(a, max(b, c)) > max(d, e) || min(a, min(b, c)) > min(d, e) {
// 		fmt.Println("NO")
// 		return
// 	}
// 	fmt.Println("YES")
// }

func metro() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)
	min1 := a*(n-1) + n
	min2 := b*(m-1) + m
	max1 := a*(n+1) + n
	max2 := b*(m+1) + m
	if min1 > max2 || min2 > max1 {
		fmt.Println("-1")
		return
	}
	//fmt.Println(max(min1, min2), min(max1, max2))
}

func sqareEquation() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	if a == 0 {
		if b == c*c && b >= 0 {
			fmt.Println("MANY SOLUTIONS")
			return
		} else {
			fmt.Println("NO SOLUTION")
			return
		}
	}
	dev := (c*c - b)
	if dev%a != 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	fmt.Println(dev / a)
}

func cleanNumbers(num string) string {
	num = strings.ReplaceAll(num, "-", "")
	num = strings.ReplaceAll(num, "(", "")
	num = strings.ReplaceAll(num, ")", "")
	num = strings.ReplaceAll(num, "+7", "8")
	if len(num) == 7 {
		num = "8495" + num
	} else if len(num) == 8 {
		num = "8495" + num[1:]
	} else if len(num) == 10 {
		num = "8" + num
	}
	return num
}

func phoneNumbers() {
	var newNum, num string
	fmt.Scan(&newNum)
	newNum = cleanNumbers(newNum)
	for i := 0; i < 3; i++ {
		fmt.Scan(&num)
		num = cleanNumbers(num)
		//fmt.Println(newNum, num)
		if newNum == num {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func Triangle() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}

func Conditioner() {
	var tRoom, tCond int
	var mode string
	fmt.Scan(&tRoom, &tCond, &mode)
	if mode == "freeze" {
		if tCond < tRoom {
			fmt.Println(tCond)
			return
		}
		fmt.Println(tRoom)
		return
	}
	if mode == "heat" {
		if tCond > tRoom {
			fmt.Println(tCond)
			return
		}
		fmt.Println(tRoom)
		return
	}
	if mode == "auto" {
		fmt.Println(tCond)
		return
	}
	fmt.Println(tRoom)
	return
}

func RLE(input string) string {
	var output string
	l := len(input)
	for i := 0; i < l; i++ {
		beg := i
		for input[i] == input[beg] {
			i++
			if i == l {
				break
			}
		}
		end := i
		i--
		cnt := end - beg
		//fmt.Println("!!!", &output)
		if cnt != 1 {
			output += fmt.Sprintf("%s%d", string(input[beg]), cnt)
		} else {
			output += fmt.Sprintf("%s", string(input[beg]))
		}
		//fmt.Println("???", &output)
	}

	return output
}
