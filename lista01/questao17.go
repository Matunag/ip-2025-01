package main
import "fmt"

func main(){
    var num1, y, i  int
    fmt.Print("Insira um número par e a quantidade de próximos números pares depois desse: ")
    fmt.Scan(&num1)
    if num1 % 2 == 0 {
        fmt.Scan(&y)
        for i=1; i<=y; i++{
            fmt.Print(num1," ")
            num1 = num1 + 2
        }
    } else{
        fmt.Println("O primeiro número não é par!")
    }
    
}    