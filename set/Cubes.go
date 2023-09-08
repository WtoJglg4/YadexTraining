package set

import (
	"fmt"
)

func Cubes() {
	var n, m int
	fmt.Scan(&n, &m)
	list1 := make([]int, n)
	list2 := make([]int, m)
	for i := range list1 {
		fmt.Scan(&list1[i])
	}
	for i := range list2 {
		fmt.Scan(&list2[i])
	}
	union := UnionTwoSetsCubes(list1, list2)
	list1 = QuickSort(list1)
	list2 = QuickSort(list2)
	fmt.Println(len(union))
	for _, v := range list1 {
		if _, ok := union[v]; ok {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n%d\n", len(list1)-len(union))
	for _, v := range list1 {
		if _, ok := union[v]; !ok {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n%d\n", len(list2)-len(union))
	for _, v := range list2 {
		if _, ok := union[v]; !ok {
			fmt.Printf("%d ", v)
		}
	}
}

func UnionTwoSetsCubes(list1, list2 []int) map[int]struct{} {

	empty1 := false
	empty2 := false
	if len(list1) == 0 {
		empty1 = true
	}

	union := make(map[int]struct{})
	//fmt.Println(union, strNumbers)
	for _, v := range list1 {
		union[v] = struct{}{}
	}
	if len(list2) == 0 {
		empty2 = true
	}
	if empty1 && empty2 {
		return make(map[int]struct{}, 0)
	}
	resultUnion := make(map[int]struct{})
	for _, v := range list2 {
		if _, ok := union[v]; ok {
			resultUnion[v] = struct{}{}
		}
	}
	for k := range union {
		delete(union, k)
	}

	return resultUnion
}
