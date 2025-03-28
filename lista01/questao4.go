package main

import "fmt"

func main(){
	var salario, energia_gasta, kW, fatura  float32
	
	fmt.Print("Insira o salário mínimo: ")
	fmt.Scan(&salario)
	fmt.Print("Insira o gasto de energia: ")
	fmt.Scan(&energia_gasta)
	kW = salario * 0.007
	fmt.Printf("O valor de cada kW é: %.2f \n", kW)
	fatura = energia_gasta * kW
	fmt.Printf("O valor da fatura é: %.2f \n", fatura)
	fatura = fatura * 0.9
	fmt.Printf("O valor da fatura com o desconto é: %.2f \n ", fatura)


}