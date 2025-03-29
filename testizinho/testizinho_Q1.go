package main 

import "fmt"

func main(){
	var num float64
	fmt.Println("Insira um número para ver se ele é negativo, positivo ou nulo")
	fmt.Scan(&num)
	if num<0 {
		fmt.Print("O número é negativo")
	} else if num == 0{
		fmt.Print("O número é nulo")
	} else {
		fmt.Print("n número é positivo")
	}
	}


