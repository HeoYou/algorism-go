package main

import (
	"bufio"
	"fmt"
	"os"
)

// output no buffer
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

// output use buffer
// func main() {
// 	start := time.Now()
// 	w := bufio.NewWriter(os.Stdout)
// 	for i := 0; i < 50000; i++ {
// 		fmt.Fprintln(w, "a")
// 	}

// 	end := time.Now()
// 	fmt.Println(end.Sub(start))
// 	//39.7799ms
// 	//3.41s
// }

// bufio 사용이 월등히 빠른걸 알 수 있음

//input use buffer 공백이 있으면 맨 앞 인풋만 들어감
func main() {
	var num int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &num)
	fmt.Println(num)
}

//input use buffer 한줄 읽기
// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	s, _ := r.ReadString('\n')
// 	fmt.Print(s)
// }
