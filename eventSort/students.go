package eventsort

import "fmt"

const (
	beg = -1
	end = 1
)

// func sortTeachers(pevents *[][]int) {
// 	events := *pevents
// 	for i := 0; i < len(events)-1; i++ {
// 		for j := i + 1; j < len(events); j++ {
// 			if events[i][0] > events[j][0] {
// 				events[i], events[j] = events[j], events[i]
// 			}
// 			if events[i][0] == events[j][0] && events[i][1] > events[j][1] {
// 				events[i][1], events[j][1] = events[j][1], events[i][1]
// 			}
// 		}
// 	}
// }

func StudentsExam() {
	var n, m int
	fmt.Scan(&n, &m)
	teachers := make([][]int, 2*m)
	for i := 0; i < 2*m; i += 2 {
		var b, e int
		fmt.Scan(&b, &e)
		teachers[i] = []int{b, beg}
		teachers[i+1] = []int{e, end}
	}
	teachers = MergeSortTeachers(teachers)
	// fmt.Println(teachers)

	withoutVision := teachers[0][0]
	onVision := 1
	for i := 1; i < len(teachers); i++ {
		// fmt.Println(i, onVision, withoutVision)
		if teachers[i][1] == beg {
			if onVision == 0 {
				withoutVision += teachers[i][0] - teachers[i-1][0] - 1
			}
			onVision++
		} else {
			onVision--
		}
	}
	withoutVision += n - 1 - teachers[len(teachers)-1][0]
	fmt.Println(withoutVision)
}

func MergeSortTeachers(list [][]int) [][]int {

	n := len(list)
	if n == 1 {
		return list
	}
	var FirstHalf, SecondHalf [][]int
	if n%2 == 0 {
		FirstHalf = list[:n/2]
		SecondHalf = list[n/2:]

	} else {
		FirstHalf = list[:n/2+1]
		SecondHalf = list[n/2+1:]
	}

	SortedFH := MergeSortTeachers(FirstHalf)
	SortedSH := MergeSortTeachers(SecondHalf)

	SortedList := MergeTeachers(SortedFH, SortedSH)
	return SortedList
}

func MergeTeachers(list_1, list_2 [][]int) [][]int {
	var SortedList [][]int

	for len(list_1) > 0 || len(list_2) > 0 {
		if len(list_1) == 0 {
			SortedList = append(SortedList, list_2...)
			return SortedList
		}
		if len(list_2) == 0 {
			SortedList = append(SortedList, list_1...)
			return SortedList
		}
		if list_1[0][0] < list_2[0][0] || list_1[0][0] == list_2[0][0] && list_1[0][1] < list_2[0][1] {
			SortedList = append(SortedList, list_1[0])
			list_1 = list_1[1:]
		} else {
			SortedList = append(SortedList, list_2[0])
			list_2 = list_2[1:]
		}
	}

	return SortedList
}
