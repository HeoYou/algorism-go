package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	visit  [][]int
	b      [][]int
	dx, dy [4]int
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)
	defer w.Flush()

	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}

	fmt.Fprint(w, dx, dy)

	var tc int
	fmt.Fscan(r, &tc)

	for t := 0; t < tc; t++ {
		var m, n, k int
		fmt.Fscan(r, &m, &n, &k)

		b = make([][]int, n)
		visit = make([][]int, n)

		for i := 0; i < n; i++ {
			b[i] = make([]int, m)
			visit[i] = make([]int, m)
		}

		for i := 0; i < k; i++ {
			var x, y int
			fmt.Fscan(r, &x, &y)

			b[y][x] = 1
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {

			}
		}

		fmt.Fprintln(w, b)
		fmt.Fprintln(w, m, n, k)
	}
}

func dfs(y, x int) {
	for i := 0; i < 4; i++ {
		xx := x + dx[i]
		yy := x + dy[i]
		if visit[yy][xx] && b[yy][xx] {

		}
	}
}
