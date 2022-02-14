package main

import "fmt"

func main() {
	var H, M int
	fmt.Scanf("%d", &H)
	fmt.Scanf("%d", &M)
	if M < 45 {
		M = M + 60 - 45
		if H == 0 {
			H = 23
		} else {
			H = H - 1
		}
	} else {
		M = M - 45
	}
	fmt.Printf("%d %d\n", H, M)
}
