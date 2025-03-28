package main

import ("fmt")
func main() {
    var horas int
    var valor int
    
    fmt.Print("escreva a quantidade de horas usadas: ")
    fmt.Scan(&horas)
    if horas%3 == 0 {
      valor = (horas / 3)  * 10
    } else if horas%3 == 1 {
      valor = ((horas - 1 ) / 3 ) * 10 + 5
    } else if horas%3 == 2 {
      valor = ((horas  - 2 ) / 3) * 10 + 10
    }
    fmt.Printf("O valor é: %.2f \n ", valor)
  }