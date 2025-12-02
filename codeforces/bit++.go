package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)

    valor := 0
    for i := 0; i < n; i++ {
        var operacao string

        fmt.Scan(&operacao)

        if operacao[1] == byte('+') {
            valor++
        } else {
            valor--
        }
    }

    fmt.Println(valor)
}
