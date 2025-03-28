package main 

import "fmt" 

func main () {

var fahrenheit, celsius float64
var quantidade int
fmt.Print("insira quantas temperaturas vão ser convertidas: ")
fmt.Scan(&quantidade)
var slice1 = []float64{}
for i := 0; i < quantidade; i++ {
    fmt.Scan(&fahrenheit)
    slice1 = append(slice1, fahrenheit )
}
for i := 0; i < quantidade; i++{
    celsius = 5 * (slice1[i] - 32) / 9
    fmt.Printf("%v Fahrenheit equivale a %.2f celsius \n", slice1[1], celsius)
}


}
