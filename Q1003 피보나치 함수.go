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

	var tc int
	fmt.Fscan(r, &tc)

	for i := 0; i < tc; i++ {
		var n int
		fmt.Fscan(r, &n)
		if n == 0 {
			fmt.Fprintln(w, 1, 0)
			continue
		}
		var a, b int = 0, 1
		for j := 0; j < n-1; j++ {
			tmp := a
			a = b
			b = a + tmp

		}
		fmt.Fprintln(w, a, b)
	}
}
