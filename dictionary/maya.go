package dictionary

import (
	"bufio"
	"fmt"
	"os"
)

func Maya() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var g, modS int
	var word, line string
	letters := make(map[rune]int)

	fmt.Fscan(in, &g, &modS, &word, &line)
	wordRune := []rune(word)
	lineRune := []rune(line)
	for _, v := range wordRune {
		letters[v]++
	}
	count := 0
	curLetters := make(map[rune]int)
	for j := 0; j < g-1; j++ {
		curLetters[lineRune[j]]++
	}
	//fmt.Println(curLetters)
	del := -1
	for i := g - 1; i < modS; i++ {
		//fmt.Println(i)
		flag := true
		curLetters[lineRune[i]]++
		del++

		if flag {
			for k, v := range letters {
				if curLetters[k] != v {
					flag = false
					break
				}
			}
			if flag {
				count++
			}

		}
		//fmt.Println(letters, curLetters)
		curLetters[lineRune[del]]--
	}
	fmt.Fprintln(out, count)
}
