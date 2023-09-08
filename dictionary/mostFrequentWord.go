package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MostFrequentWord() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	wordsCount := make(map[string]int)
	var words []string

	for {
		line, _ := in.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if line == "" || len(line) < 2 {
			if len(words) < 1 {
				return
			}
			break
		}
		words = strings.Fields(line)
		for _, v := range words {
			wordsCount[v]++
		}
		//fmt.Printf("%v, %v, %d", words, wordsCount, len(words))
		if len(words) < 1 {
			return
		}
	}

	max := wordsCount[words[0]]
	key := []rune(words[0])
	for k, v := range wordsCount {
		if v > max {
			max = v
			key = []rune(k)
		}
		if v == max {
			kRune := []rune(k)
			for j := 0; j < min(len(kRune), len(key)); j++ {
				if kRune[j] < key[j] {
					key = kRune
					break
				}
				if kRune[j] > key[j] {
					break
				}
				if j == min(len(kRune), len(key))-1 && key[j] == kRune[j] {
					if len(kRune) < len(key) {
						key = kRune
						break
					}
				}
			}
		}
	}
	fmt.Fprintln(out, string(key))
}
