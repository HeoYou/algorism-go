package main

import (
    "fmt"
    "os"
    "bufio"
)

func main(){
    w := bufio.NewWriter(os.Stdout)
    r := bufio.NewReader(os.Stdin)
    defer w.Flush()

    var num int
    fmt.Fscan(r, &num)

    if num % 2 == 0{
        fmt.Fprint(w, "CY")
    } else{
        fmt.Fprint(w, "SK")
    }
}
