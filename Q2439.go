package main

import "fmt"

func main() {
	var a int
	star := "*"
	fmt.Scanln(&a)

	for i := a; i > 0; i-- {

		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}

		fmt.Println(star)
		star += "*"
	}
}
