package main
import "fmt"

func main(){
    var valor_inicial, razão, n, i, soma  float32
    fmt.Print("Escreva o valor inicial, a razão e a quantidade de números: ")
    fmt.Scan(&valor_inicial, &razão, &n)
    for i=0; i<n; i++ {
        soma += valor_inicial + razão * i
    }
    fmt.Println("O valor da soma total é:", soma)
}   