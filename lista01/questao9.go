package main 

import "fmt" 

func main () {

var a, b, c, delta float64
fmt.Print("insira os valores de a,b, c: ")
fmt.Scan(&a, &b, &c)
delta = b * b - 4 * a * c
fmt.Printf("O valor de delta é = %.2f \n", delta)

}
