package main

import "fmt"

func main() {
    var casos int

    fmt.Scan(&casos)

    for i := 0; i < casos; i++ {
        var n, k int64

        fmt.Scan(&n, &k)

        var esquerda int64 = 1
        var direita int64 = 2 * k
        var resposta int64 = 0

        for esquerda <= direita {
            meio := (esquerda + direita) / 2

            qtdNaoDivisel := meio - (meio / n)

            if qtdNaoDivisel >= k {
                resposta = meio
                direita = meio - 1
            } else {
                esquerda = meio + 1
            }
        }

        fmt.Println(resposta)
    }
}
