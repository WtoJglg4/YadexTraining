package set

type Set [][]int

func NewSet(size int) Set {
	set := make([][]int, size)
	for i := range set {
		set[i] = make([]int, 0)
	}

	return set
}

func (s Set) Add(el int) {
	s[el%len(s)] = append(s[el%len(s)], el)
}

func (s Set) Find(el int) bool {
	for _, v := range s[el%len(s)] {
		if el == v {
			return true
		}
	}
	return false
}
