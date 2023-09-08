package set

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Text() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// words = append(words, strings.Split(text, `\s+`)...)
		words = append(words, regexp.MustCompile(`\s+`).Split(strings.TrimSpace(text), -1)...)

	}

	unique := make(map[string]struct{})
	for _, w := range words {
		unique[w] = struct{}{}
	}
	//words := make(map[string]struct{})
	// words := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(str), -1)

	// fmt.Println(words)
	// fmt.Println(unique)
	fmt.Println(len(unique))
}
