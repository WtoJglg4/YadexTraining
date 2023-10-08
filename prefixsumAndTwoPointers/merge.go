package prefixsumAndTwoPointers

func Merge(list1, list2 []int) []int {
	resList := make([]int, len(list1)+len(list2))
	index1, index2 := 0, 0
	for i := range resList {
		if index1 < len(list1) && (index2 == len(list2) || list1[index1] <= list2[index2]) {
			resList[i] = list1[index1]
			index1++
		} else {
			resList[i] = list2[index2]
			index2++
		}
	}
	return resList
}

func MergeTwo(list1, list2, list3, list4 []int) ([]int, []int) {
	resList := make([]int, len(list1)+len(list2))
	resList2 := make([]int, len(list1)+len(list2))
	index1, index2 := 0, 0
	for i := range resList {
		if index1 < len(list1) && (index2 == len(list2) || list1[index1] <= list2[index2]) {
			resList[i] = list1[index1]
			resList2[i] = list3[index1]
			index1++
		} else {
			resList[i] = list2[index2]
			resList2[i] = list4[index2]
			index2++
		}
	}
	return resList, resList2
}
