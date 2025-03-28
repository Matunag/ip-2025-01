package main

import ("fmt")
func main() {
    var nota float32
    
    fmt.Print("escreva a nota a ser analisada: ")
    fmt.Scan(&nota)
    if nota<=6 {
        fmt.Printf("Sua nota é: %.1f e seu conceito é D \n", nota)
        } else if nota>=6 && nota<7.5{
            fmt.Printf("Sua nota é: %.1f e seu conceito é C \n", nota)
                } else if nota>=7.5 && nota<9{
                fmt.Printf("Sua nota é: %.1f e seu conceito é B \n", nota)
                } else{
                fmt.Printf("Sua nota é: %.1f e seu conceito é A \n", nota)
                }
    
    }