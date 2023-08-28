package linear

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProductThree2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strNums := strings.Split(input, " ")
	seq := make([]int, len(strNums))
	for i, v := range strNums {
		seq[i], _ = strconv.Atoi(v)
	}

	max1, max2, max3 := seq[0], seq[1], seq[2]
	i1, i2, i3 := 0, 1, 2
	sortThreeMax(&max1, &max2, &max3, &i1, &i2, &i3)
	fmt.Println(max1, max2, max3, i1, i2, i3)
	for i := 3; i < len(seq); i++ {
		v := seq[i]
		if v > max1 {
			max3, i3 = max2, i2
			max2, i2 = max1, i1
			max1, i1 = v, i
		} else if v > max2 {

		}
	}
}
func sortThreeMax(pm1, pm2, pm3, pi1, pi2, pi3 *int) {
	m1, m2, m3, i1, i2, i3 := *pm1, *pm2, *pm3, *pi1, *pi2, *pi3
	if m1 > m2 && m1 > m3 {
		if m2 > m3 {
			return
		} else {
			m2, m3 = m3, m2
			i2, i3 = i3, i2
		}
	} else if m2 > m1 && m2 > m3 {
		if m1 < m3 {
			m1, m3 = m3, m1
			i1, i3 = i3, i1

		}
		m1, m2 = m2, m1
		i1, i2 = i2, i1
		return
	} else {
		if m2 < m1 {
			m1, m2 = m2, m1
			i1, i2 = i2, i1
		}
		m1, m3 = m3, m1
		i1, i3 = i3, i1
		return
	}
}

func maxProductThree() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strNums := strings.Split(input, " ")
	seq := make([]int, len(strNums))
	for i, v := range strNums {
		seq[i], _ = strconv.Atoi(v)
	}
	//fmt.Println(nums)
	l := len(seq)
	rearrange(&seq, 0, l, l-1)
	//fmt.Println(seq)
	rearrange(&seq, 0, l-1, l-2)
	//fmt.Println(seq)
	rearrange(&seq, 0, l-3, 2)
	//fmt.Println(seq)
	prod1, prod2 := seq[l-1]*seq[l-2]*seq[l-3], seq[0]*seq[1]*seq[l-1]
	if prod1 > prod2 {
		fmt.Println(seq[l-1], seq[l-2], seq[l-3])
	} else {
		fmt.Println(seq[l-1], seq[0], seq[1])
	}

}

func rearrange(pseq *[]int, left, right, k int) {
	seq := *pseq
	for left < right {
		gtxFirst, eqxFirst := left, left
		x := seq[(left+right)/2]
		for i := left; i < right; i++ {
			now := seq[i]
			if now == x {
				seq[i] = seq[gtxFirst]
				seq[gtxFirst] = now
				gtxFirst++
			} else if now < x {
				seq[i] = seq[gtxFirst]
				seq[gtxFirst] = seq[eqxFirst]
				seq[eqxFirst] = now
				eqxFirst++
				gtxFirst++
			}
		}
		if k < eqxFirst {
			right = eqxFirst - 1
		} else if k >= gtxFirst {
			left = gtxFirst
		} else {
			return
		}
	}

}
