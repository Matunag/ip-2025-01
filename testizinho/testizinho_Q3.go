package main

import "fmt"
func main(){
soma := 0

	for i := 1 ; i <= 100; i++{
		fmt.Print(i, ", ")
		soma += i
	}
	fmt.Print("O valor da soma é: ", soma)
}