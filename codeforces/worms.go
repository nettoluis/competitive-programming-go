package main

import "fmt"

func buscaBinaria(sequencia []int, valor int) int {
    direita := len(sequencia) - 1
    esquerda := 0
    resposta := -1

    for direita >= esquerda {
        meio := esquerda + (direita - esquerda) / 2

        if sequencia[meio] >= valor {
            resposta = meio

            direita = meio - 1
        } else {
            esquerda = meio + 1
        }
    }

    return resposta + 1
}

func main() {
    var pilhas int

    fmt.Scan(&pilhas)
    prefixSums := make([]int, pilhas)
    fmt.Scan(&prefixSums[0])

    for i := 1; i < pilhas; i++ {
        var valorMaximoPilha int

        fmt.Scan(&valorMaximoPilha)

        prefixSums[i] = prefixSums[i-1] + valorMaximoPilha
    }

    var vermes int

    fmt.Scan(&vermes)

    for j := 0; j < vermes; j++ {
        var saborosa int

        fmt.Scan(&saborosa)

        fmt.Println(buscaBinaria(prefixSums, saborosa))
    }
}
