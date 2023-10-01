package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"github/YandexTraining/mathem"
	"os"
)

func StylishClothes() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n)
	shirts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &shirts[i])
	}
	fmt.Fscan(in, &m)
	trousers := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &trousers[i])
	}

	minDiff := mathem.Abs(shirts[0] - trousers[0])
	currDiff, sIndex, tIndex, minsIndex, mintIndex := 0, 0, 0, 0, 0

	for {

		currDiff = mathem.Abs(shirts[sIndex] - trousers[tIndex])
		if currDiff < minDiff {
			minDiff = currDiff
			minsIndex = sIndex
			mintIndex = tIndex
		}

		if minDiff == 0 {
			break
		}

		if sIndex == len(shirts) || trousers[tIndex] < shirts[sIndex] {
			tIndex++
			if tIndex == len(trousers) {
				tIndex--
				break
			}
		} else {
			sIndex++
			if sIndex == len(shirts) {
				sIndex--
				break
			}
		}
	}
	fmt.Fprint(out, shirts[minsIndex], trousers[mintIndex])
}
