package testing

import (
	"fmt"
	"github/YandexTraining/prefixsumAndTwoPointers"
	"math/rand"
	"time"
)

// const letterBytes = "abcdefghijklmnopqrstuvwxy"

func Testing() {
	n := 10

	str1 := make([]prefixsumAndTwoPointers.Point, n)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		for i := 0; i < len(str1); i++ {
			str1[i].X = r1.Intn(10)
			str1[i].Y = r1.Intn(10)
		}
		fmt.Println(str1)
		ans1, ans2 := prefixsumAndTwoPointers.TrianglesTest(str1), prefixsumAndTwoPointers.TrianglesSlow(str1)
		if ans1 == ans2 {
			fmt.Printf("%v\n%v %v\n", str1, ans1, ans2)
		} else {
			fmt.Printf("ERROR!!! %v\n%v %v\n", str1, ans1, ans2)
			break
		}
	}
}
