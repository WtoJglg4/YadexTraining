package dictionary

import "fmt"

func Synonims() {
	var n int
	fmt.Scan(&n)
	dict := make(map[string]string)
	for i := 0; i < n; i++ {
		var word1, word2 string
		fmt.Scan(&word1, &word2)
		dict[word1] = word2
		dict[word2] = word1
	}
	var word string

	fmt.Scan(&word)
	fmt.Println(dict[word])
}
