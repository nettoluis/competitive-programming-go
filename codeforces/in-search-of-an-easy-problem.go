package main

import "fmt"

func main() {
    var qtd int

    fmt.Scan(&qtd)

    valor := 0
    for i := 0; i < qtd; i++ {
        fmt.Scan(&valor)

        if valor != 0 {
            fmt.Println("HARD")
            return
        }
    }

    fmt.Println("EASY")
}
