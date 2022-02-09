package main

import "fmt"

func main() {
	var s int
	fmt.Scanln(&s)

	s = s / 10

	if s > 8 {
		fmt.Println("A")
	} else if s > 7 {
		fmt.Println("B")
	} else if s > 6 {
		fmt.Println("C")
	} else if s > 5 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
