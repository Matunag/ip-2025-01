package main

import("fmt")

func main() {
    
    var nota1, nota2, nota3 float32
    fmt.Println("insira as três notas")
    fmt.Scan(&nota1)
    fmt.Scan(&nota2)
    fmt.Scan(&nota3)
    var media float32 = (nota1 + nota2 + nota3) / 3
    fmt.Printf ("A média é %.2f \n:", media) 
    if (media>=6) {
        fmt.Println("APROVADO")
       } else {
        fmt.Println("REPROVADO")
    }
    
}