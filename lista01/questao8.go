package main 

import "fmt" 

func main () {

var raio, altura, area, preco float64
const pi float64 = 3.14159
fmt.Print("insira os valores de raio e altura da lata: ")
fmt.Scan(&raio, &altura)
area = 2 * (pi * raio * raio) + 2 * pi * raio * altura
preco = area * 100
fmt.Printf("O valor do custo é = %.2f \n", preco)

}
