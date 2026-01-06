package main

import "fmt"

func main() {
    var k, dinheiroInicial, qtd int64

    fmt.Scan(&k, &dinheiroInicial, &qtd)

    custoTotal := ((k + (k * qtd)) * qtd / 2) 

    switch {
        case custoTotal > dinheiroInicial:
            fmt.Println(custoTotal - dinheiroInicial)
        default:
            fmt.Println(0)
    }
}
