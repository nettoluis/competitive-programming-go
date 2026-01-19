package main

import (
    "fmt"
)

func merge(a, b []float64) []float64 {
    c := make([]float64, len(a) + len(b))
    idC := 0

    for len(a) != 0 && len(b) != 0 {
        if a[0] > b[0] {
            c[idC] = b[0]
            b = b[1:]
        } else {
            c[idC] = a[0]
            a = a[1:]
        }
            idC++
    }

    for len(a) != 0 {
        c[idC] = a[0]
        a = a[1:]
        idC++
    }

    for len(b) != 0 {
        c[idC] = b[0]
        b = b[1:]
        idC++
    }

    return c
}

func mergeSort(a[]float64) []float64 {
    if len(a) <= 1 {
        return a
    }

    primeiraMetade := a[:len(a)/2]
    segundaMetade := a[len(a)/2:]

    primeiraMetade = mergeSort(primeiraMetade)
    segundaMetade = mergeSort(segundaMetade)

    return merge(primeiraMetade, segundaMetade)
}

func maiorDiferenca(a []float64) float64 {
    var maiorDiferenca float64 = 0
    for i := 1; i < len(a); i++ {
        if (a[i] - a[i - 1]) > maiorDiferenca {
            maiorDiferenca = a[i] - a[i - 1]
        }
    }

    return maiorDiferenca
}

func main() {
    var qtdLanternas int
    var tamanhoRua float64

    fmt.Scan(&qtdLanternas, &tamanhoRua)

    posicoes := make([]float64, qtdLanternas)
    for i := 0; i < len(posicoes); i++ {
        fmt.Scan(&posicoes[i])
    }

    posicoes = mergeSort(posicoes)

    maiorDistancia := maiorDiferenca(posicoes)

    fmt.Printf("%.10f\n", max(maiorDistancia/2, max(posicoes[0] - 0, tamanhoRua - posicoes[qtdLanternas - 1])))
}
