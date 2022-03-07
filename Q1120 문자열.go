package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var str1, str2 string

	fmt.Fscan(r, &str1)
	fmt.Fscan(r, &str2)

	minData := 100
	var count int
	for i := 0; i < len(str2)-len(str1)+1; i++ {
		count = 0
		for j := 0; j < len(str1); j++ {
			if str1[j] != str2[j+i] {
				count += 1
			}
		}
		if count < minData {
			minData = count
		}
	}
	fmt.Fprintln(w, minData)
}
