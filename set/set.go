package set

type Set [][]int

func NewSet(setSize int) Set {
	set := make([][]int, setSize)
	for i := range set {
		set[i] = make([]int, 0)
	}

	return set
}

func (s Set) Add(el int) {
	setSize := len(s)
	s[el%setSize] = append(s[el%setSize], el)
}

func (s Set) Find(el int) bool {
	setSize := len(s)
	for _, v := range s[el%setSize] {
		if el == v {
			return true
		}
	}
	return false
}

func (ps *Set) Delete(el int) {
	s := *ps
	setSize := len(s)
	if s.Find(el) {
		for i, v := range s[el%setSize] {
			if el == v {
				newS := s[el%setSize][:i]
				newS = append(newS, s[el%setSize][i+1:]...)
				s[el%setSize] = newS
				return
			}
		}
	}
}
