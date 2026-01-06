package main

import "fmt"

func main() {
    var qtd int
    var sequencia string

    fmt.Scan(&qtd, &sequencia)

    contador := 0
    for i := 0; i < (qtd - 1); i++ {
        if sequencia[i] == sequencia[i + 1] {
            contador++
        }
    }

    fmt.Println(contador)


}
