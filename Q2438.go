package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 1; i < num+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
