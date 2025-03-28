package main

import "fmt"

func main (){
	var agua, fatura float32
	var conta int
	var categoria string
	fmt.Print("Conta do cliente, gasto de água e categoria (R para residencial, I para industrial e C para comercial): ")
	fmt.Scan(&conta, &agua, &categoria)
	switch categoria {
	case "R", "r":
		fatura = 5 + 0.05 * agua
	case "I", "i":
		if agua>100{
			fatura = 800 + 0.04 * (agua - 100)
		} else {
			fatura = 800
		}
	case "C", "c":
	if agua>80{
		fatura = 500 + 0.25 * (agua - 80)
	} else {
		fatura = 500
	}

	}
	fmt.Printf("A conta é: %v e a fatura é de: R$ %.2f \n", conta, fatura)
}