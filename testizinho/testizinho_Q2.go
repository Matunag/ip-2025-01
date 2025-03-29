package main

import "fmt"

func main (){
	fmt.Println("Insira o número a ser analisado:")
	var num float64
	fmt.Scan(&num)
if num <90 && num>20 {
	fmt.Print("O número está entre 20 e 90")
} else{
	fmt.Print("o número não está entre 20 e 90")
}
}