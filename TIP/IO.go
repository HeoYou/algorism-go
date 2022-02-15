package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 50000; i++ {
		fmt.Println("abc")
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
}
