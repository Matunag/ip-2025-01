package main

import "fmt"

func main () {
	var horas, minutos, segundos, total_em_segundos int32
	fmt.Println("Insira a quantidade de horas:")
	fmt.Scan(&horas)
	fmt.Println("insira a quantidade de minutos")
	fmt.Scan(&minutos)
	fmt.Println("insira a quantidade de segundos")
	fmt.Scan(&segundos)
	total_em_segundos = horas * 3600 + minutos * 60 + segundos
	fmt.Println("o total de tempo em segundos é:", total_em_segundos)
	
}