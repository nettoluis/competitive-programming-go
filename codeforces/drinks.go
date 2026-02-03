package main

import "fmt"

func main() {
    var qtd uint64

    fmt.Scan(&qtd)

    var encontrado float64 = 0
    for i := 0; i < int(qtd); i++ {
        var porcentagem float64

        fmt.Scan(&porcentagem)

        encontrado += porcentagem
    }

    fmt.Printf("%.12f\n", encontrado / float64(qtd))
}
