package main

import "fmt"

func main() {
	var num int
	fmt.Println("Escreva o número a ser analisado: ")
	fmt.Scanln(&num)
	if num % 3 == 0 && num % 5 == 0 {
		fmt.Print("O número é divisível!!!")
	} else { 
		fmt.Print("O número não é divisível!!!")
	}
}
