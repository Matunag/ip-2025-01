package main 

import "fmt" 

func main () {

var fahrenheit, celsius, polegadas, milimetros float64
fmt.Print("insira os valores de farenheit e polegadas a serem convertidos: ")
fmt.Scan(&fahrenheit, &polegadas)
celsius = 5 * (fahrenheit - 32) / 9
milimetros = polegadas * 25.4
fmt.Printf("O valor em celsius = %.2f \n", celsius)
fmt.Printf("O valor em milimetros = %.2f", milimetros)

}
