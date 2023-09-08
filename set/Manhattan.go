package set

import "fmt"

type Point struct {
	x, y int
}

func Distance(point1, point2 Point) int {
	return abs(point1.x-point2.x) + abs(point1.y-point2.y)
}

func (p Point) AreaGeneration(t, d int) map[Point]bool {
	Area := make(map[Point]bool)
	d = min(t, d)
	for x := p.x - d; x <= p.x+d; x++ {
		for y := p.y - d; y <= p.y+d; y++ {
			point := Point{x, y}
			if Distance(p, point) <= d {
				Area[point] = false
			}
		}
	}
	return Area
}

func Manhattan() {
	var t, d, n int
	var lastPoint, preLastPoint Point
	fmt.Scan(&t, &d, &n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		if i == n-2 {
			preLastPoint.x = x
			preLastPoint.y = y
		}
		if i == n-1 {
			lastPoint.x = x
			lastPoint.y = y
		}
	}
	lastArea := lastPoint.AreaGeneration(t, d)
	preLastArea := preLastPoint.AreaGeneration(t, d)

	count := 0
	for p1, _ := range lastArea {
		for p2, _ := range preLastArea {
			if Distance(p1, p2) <= t {
				lastArea[p1] = true
				count++
				break
			}
		}
	}
	fmt.Println(count)
	for p1, v := range lastArea {
		if v {
			fmt.Println(p1.x, p1.y)
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
