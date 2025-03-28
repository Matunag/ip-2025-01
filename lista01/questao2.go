package main 

import "fmt"

func main (){
var pessoas, popular, geral, arquibancada, cadeiras, arrecadacao, jogos int

fmt.Print("Insira a quantidade de jogos: ")
fmt.Scan(&jogos)
for i := 1 ; i<=jogos; i++ {
	fmt.Println("Escreva a quantidade de pessoas: ")
	fmt.Scan(&pessoas)
	fmt.Println("Escreva a porcentagem de pessoas que são da categoria popular: ")
	fmt.Scan(&popular)
	fmt.Println("Escreva a porcentagem de pessoas que são da categoria geral: ")
	fmt.Scan(&geral)
	fmt.Println("Escreva a porcentagem de pessoas que são da categoria arquibancada: ")
	fmt.Scan(&arquibancada)
	fmt.Println("Escreva a porcentagem de pessoas que são da categoria cadeiras: ")
	fmt.Scan(&cadeiras)
	arrecadacao = pessoas * (popular + geral * 5 + arquibancada * 10 + cadeiras * 20) / 100
	fmt.Printf("A arrecadação do jogo N. %v foi %.2f \n", i, arrecadacao)
}






}