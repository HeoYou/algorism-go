package main

import "fmt"

func main() {
	var y int
	fmt.Scanln(&y)

	if y%4 == 0 {
		if y%400 == 0 {
			fmt.Println(1)
		} else if y%100 != 0 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	} else {
		fmt.Println(0)
	}
}
