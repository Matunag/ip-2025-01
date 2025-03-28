package main
import "fmt"

func main(){
    var n, i float32
    var soma float32 = 0
    fmt.Print("escreva o valor de n:")
    fmt.Scan(&n)
    if n<=1 {
        fmt.Print("Número inválido:")
    } else{
        for i=1; i<=n; i++ {
        soma += 1 / i
        }
     fmt.Println("o valor da soma é: %.6", soma)
        
    }
}