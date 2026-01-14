package main

import "fmt"

func somaNumeros(numeros []int) int {
    contador := 0
    for _, v := range numeros {
        contador += v
    }

    return contador
}

func main() {
    var n int

    fmt.Scan(&n)
    numeros := make([]int, n - 1)

    for i := 0; i < len(numeros); i++ {
        fmt.Scan(&numeros[i])
    }

    somaEncontrada := somaNumeros(numeros)
    somaEsperada := ((1 + n) * n) / 2

    fmt.Println(somaEsperada - somaEncontrada)
}
