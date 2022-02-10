// package main

// import "fmt"

// func main() {
// 	var a int
// 	star := "*"
// 	fmt.Scanln(&a)

// 	for i := a; i > 0; i-- {

// 		for j := 1; j < i; j++ {
// 			fmt.Print(" ")
// 		}

// 		fmt.Println(star)
// 		star += "*"
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var a int
	fmt.Fscanln(r, &a)
	fmt.Println(a)

	for i := a; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Fprintf(w, " ")
			fmt.Print("123")
		}
		for j := 0; j < i; j++ {
			fmt.Fprintf(w, "*")
		}
	}
}
