package main
import "fmt"

func main(){
    var salario float32
    fmt.Print("Insira seu salário: ")
    fmt.Scan(&salario)
    if salario<=300 {
        salario *= 1.5
    } else{
        salario *= 1.3
    }
    fmt.Printf("Seu salário após o reajuste é de: %.2f \n ", salario)
}    