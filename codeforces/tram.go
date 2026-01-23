package main

import "fmt"

func main() {
    var paradas int

    fmt.Scan(&paradas)

    capacidadeMaxima, capacidadeAtual := 0, 0
    for i := 0; i < paradas; i++ {
        var saidas, entradas int

        fmt.Scan(&saidas, &entradas)

        capacidadeAtual += (entradas - saidas)

        if capacidadeAtual > capacidadeMaxima {
            capacidadeMaxima = capacidadeAtual
        }
    }

    fmt.Println(capacidadeMaxima)
}
