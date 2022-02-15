package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// func main() {
// 	start := time.Now()
// 	for i := 0; i < 50000; i++ {
// 		fmt.Print("abc")
// 	}

// 	end := time.Now()
// 	fmt.Println(end.Sub(start))
// 	//5.8s ln
// 	//1.49s
// }

func main() {
	start := time.Now()
	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < 50000; i++ {
		fmt.Fprintln(w, "a")
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
	//ln 39.7799ms
}
