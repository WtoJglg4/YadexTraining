package dictionary

func CountSort(pseq *[]int) {
	seq := *pseq
	if len(seq) < 2 {
		return
	}
	minEl, maxEl := seq[0], seq[0]
	for _, v := range seq {
		if v < minEl {
			minEl = v
		}
		if v > maxEl {
			maxEl = v
		}
	}
	k := maxEl - minEl
	count := make([]int, k+1)
	for _, v := range seq {
		count[v-minEl]++
	}
	seq2 := make([]int, 0)
	for i, v := range count {
		for j := 0; j < v; j++ {
			seq2 = append(seq2, i+minEl)
		}
	}
	*pseq = seq2
}
