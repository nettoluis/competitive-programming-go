package main

import "fmt"

func main() {
    var tamanho, lugar int

    fmt.Scan(&tamanho, &lugar)

    pontuacoes := make([]int, tamanho)

    for i := 0; i < tamanho; i ++ {
        fmt.Scan(&pontuacoes[i])
    }

    recorde := pontuacoes[lugar - 1]

    passados := 0
    for i := 0; i < tamanho; i++ {
        if pontuacoes[i] >= recorde && pontuacoes[i] > 0 {
            passados++
        }
    }

    fmt.Println(passados)
}
