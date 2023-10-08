package prefixsumAndTwoPointers

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	X, Y int
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(p1.X-p2.X)), 2) + math.Pow(math.Abs(float64(p1.Y-p2.Y)), 2))
}

func C(k, n int) int {
	res := (Fact(n-k) * Fact(k))
	if res == 0 {
		fmt.Printf("\n!!! C: k: %v ,n: %v\n", k, n)
	}
	return Fact(n) / res
}

func Triangles() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	points := make([]Point, n)
	for i := range points {
		fmt.Fscan(in, &points[i].X, &points[i].Y)
	}

	count := 0
	for i := 0; i < n; i++ {
		distances := make(map[float64][]Point)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dist := distance(points[i], points[j])
			if _, ok := distances[dist]; !ok {
				distances[dist] = make([]Point, 0)
			}
			distances[dist] = append(distances[dist], points[j])
		}
		for d, v := range distances {
			k := 1
			for p := 0; p < len(v)-1; p++ {
				for m := p + 1; m < len(v); m++ {
					if distance(v[p], v[m]) != 2*d {
						k++
						// fmt.Fprintln(out, k, v[p], v[m], distance(v[p], v[m]))
					}
				}
			}
			if k > 1 {
				count += C(2, k)
			}

		}
		// fmt.Fprintln(out, count, distances)
	}
	fmt.Fprintln(out, count)
}

func TrianglesTest(points []Point) int {
	n := len(points)

	count := 0
	for i := 0; i < n; i++ {
		distances := make(map[float64][]Point)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dist := distance(points[i], points[j])
			if _, ok := distances[dist]; !ok {
				distances[dist] = make([]Point, 0)
			}
			distances[dist] = append(distances[dist], points[j])
		}
		for d, v := range distances {
			k := 1
			for p := 0; p < len(v)-1; p++ {
				for m := p + 1; m < len(v); m++ {
					if distance(v[p], v[m]) != 2*d {
						k++
						// fmt.Fprintln(out, k, v[p], v[m], distance(v[p], v[m]))
					}
				}
			}
			if k > 1 {
				count += C(2, k)
			}

		}
		// fmt.Fprintln(out, count, distances)
	}
	return count
}

func TrianglesSlow(points []Point) int {
	n := len(points)
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if i == j || i == k || j == k {
					continue
				}

				a := distance(points[i], points[j])
				b := distance(points[j], points[k])
				c := distance(points[i], points[k])
				if S(a, b, c) > 0 && (a == b || a == c || b == c) {
					count++
				}
			}
		}
	}
	return count
}

func S(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}
