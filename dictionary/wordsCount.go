package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordsCount() {
	reader := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	words := make(map[string]int)
	for {

		str, _ := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		//fmt.Println(str, len(str))
		if str == "" || len(str) == 1 {
			break
		}

		strs := strings.Fields(str)
		for _, v := range strs {
			fmt.Fprintf(out, "%d ", words[v])
			words[v]++
		}
	}
}
