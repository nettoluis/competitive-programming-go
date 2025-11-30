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

    fmt.Scanf("%d\n", &a)
    fmt.Scanf("%d\n", &b)

    x := somaRecursiva(a, b)

    fmt.Printf("X = %d\n", x)
}
