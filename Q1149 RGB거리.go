package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)

	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)

	var answer [][]int

	var input []int
	var cr, cg, cb int
	fmt.Fscan(r, &cr, &cg, &cb)
	input = append(input, cr, cg, cb)
	answer = append(answer, input)
	for i := 1; i < n; i++ {
		var input []int
		var cr, cg, cb int
		fmt.Fscan(r, &cr, &cg, &cb)

		inputNum := Min(answer[i-1][1], answer[i-1][2])
		input = append(input, cr+inputNum)

		inputNum = Min(answer[i-1][0], answer[i-1][2])
		input = append(input, cg+inputNum)

		inputNum = Min(answer[i-1][0], answer[i-1][1])
		input = append(input, cb+inputNum)

		answer = append(answer, input)
	}

	minAnswer := Min(answer[n-1][0], answer[n-1][1])
	minAnswer = Min(minAnswer, answer[n-1][2])

	fmt.Fprintln(w, minAnswer)

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
