package main
import "fmt"

func main(){
	var a, b, c, num, quadrado int
	fmt.Print("Insira os três algarismos: ")
	fmt.Scan(&a, &b, &c)
	num = 100 * a + 10 * b + c
	quadrado = num * num
	fmt.Printf("%v, %v", num, quadrado)
}