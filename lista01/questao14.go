package main

import ("fmt"; "math")
func main() {
    var h, a, volume float64
    
    fmt.Print("escreva a altura e a aresta da base em metros: ")
    fmt.Scan(&h, &a)
    var raiz_de_tres float64 = math.Sqrt(3)
    volume = (3 * a * a * raiz_de_tres ) * h / 6
    fmt.Printf("O volume da pirâmide é: %.2f", volume)   
    
}