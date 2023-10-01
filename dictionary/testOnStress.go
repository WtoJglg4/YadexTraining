package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func TestOnStress() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	dict := make(map[string][]int)
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)
		for i, v := range []rune(word) {
			if unicode.IsUpper(v) {
				dict[strings.ToLower(word)] = append(dict[strings.ToLower(word)], i)
				break
			}
		}
	}

	smth, _ := in.ReadString('\n')
	smth += "a"
	line, _ := in.ReadString('\n')
	words := strings.Fields(line)

	mistakes := 0
	for _, w := range words {
		upperCount, upperIndex := 0, 0
		for j, sym := range []rune(w) {
			if unicode.IsUpper(sym) {
				upperCount++
				upperIndex = j
			}
			if upperCount > 1 {
				break
			}
		}
		if upperCount != 1 {
			mistakes++
			continue
		}

		if _, ok := dict[strings.ToLower(w)]; !ok {
			continue
		}

		inDict := false
		for _, v := range dict[strings.ToLower(w)] {
			if v == upperIndex {
				inDict = true
				break
			}
		}
		if !inDict {
			mistakes++
			continue
		}

	}

	fmt.Fprintln(out, mistakes)

}
