package main

import "fmt"

func main() {
    var N, i, x int
    
    fmt.Print("escreva o valor de N: ")
    fmt.Scan(&N)
    
    for i=0; i<=N; i++{
        if i % 2 == 0 {
            x = i * i
        fmt.Println(x)
        }
    }
}