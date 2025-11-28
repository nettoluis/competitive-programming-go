package main

import "fmt"

func somaRecursiva(a, b int) int {
    if b == 0 {
        return a
    } else {
        if b < 0 {
            return somaRecursiva(a - 1, b + 1)
        } else {
            return somaRecursiva(a + 1, b - 1)
        }
    }
}

func main() {
    var a, b int

    fmt.Scan(&a, &b)

    soma := somaRecursiva(a, b)

    fmt.Printf("SOMA = %d\n", soma)
}
