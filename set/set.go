package set

type Set [][]int

func NewSet(setSize int) Set {
	set := make([][]int, setSize)
	for i := range set {
		set[i] = make([]int, 0)
	}
	return set
}

func (s Set) Add(els ...int) {
	for _, el := range els {
		if !s.Find(el) {
			setSize := len(s)
			bucket := abs(el % setSize)
			s[bucket] = append(s[bucket], el)
		}
	}

}

func (s Set) Find(el int) bool {
	setSize := len(s)
	bucket := abs(el % setSize)
	for _, v := range s[bucket] {
		if el == v {
			return true
		}
	}
	return false
}

func (ps *Set) Delete(el int) {
	s := *ps
	setSize := len(s)
	bucket := abs(el % setSize)
	if s.Find(el) {
		for i, v := range s[bucket] {
			if el == v {
				newS := s[bucket][:i]
				newS = append(newS, s[bucket][i+1:]...)
				s[bucket] = newS
				return
			}
		}
	}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
