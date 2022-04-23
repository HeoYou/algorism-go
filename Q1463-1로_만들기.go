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
	dp := make([]int, 1000001)
	var n int64
	fmt.Fscan(r, &n)
	for n > 1 {
		if dp[n-1] > dp[n]+1 || dp[n-1] == 0 {
			dp[n-1] = dp[n] + 1
		}
		if n%2 == 0 && (dp[n/2] > dp[n]+1 || dp[n/2] == 0) {
			dp[n/2] = dp[n] + 1
		}
		if n%3 == 0 && (dp[n/3] > dp[n]+1 || dp[n/3] == 0) {
			dp[n/3] = dp[n] + 1
		}
		n -= 1
	}

	fmt.Fprint(w, dp[1])
}
