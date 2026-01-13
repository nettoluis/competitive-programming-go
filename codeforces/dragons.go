package main

import (
    "fmt"
)

func merge(a, b [][2]int) [][2]int {
    c := make([][2]int, len(a) + len(b))
    idC := 0

    for len(a) != 0 && len(b) != 0 {
        if a[0][0] > b[0][0] {
            c[idC][0], c[idC][1] = b[0][0], b[0][1]
            b = b[1:]
        } else {
            c[idC][0], c[idC][1] = a[0][0], a[0][1]
            a = a[1:]
        }
        idC++
    }

    for len(a) != 0 {
            c[idC][0], c[idC][1] = a[0][0], a[0][1]
            a = a[1:]
            idC++
    }

    for len(b) != 0 {
            c[idC][0], c[idC][1] = b[0][0], b[0][1]
            b = b[1:]
            idC++
    }

    return c
}

func mergesort(a [][2]int) [][2]int {
    if len(a) == 1 {
        return a
    }

    primeiraMetade := a[:(len(a)/2)]
    segundaMetade := a[(len(a)/2):]

    primeiraMetade = mergesort(primeiraMetade)
    segundaMetade = mergesort(segundaMetade)

    return merge(primeiraMetade, segundaMetade)
}

func main() {
    var forcaKirito, qtdInimigos int

    fmt.Scan(&forcaKirito, &qtdInimigos)

    inimigos := make([][2]int, qtdInimigos)
    
    for i := 0; i < qtdInimigos; i++ {
        var forca, bonus int

        fmt.Scan(&forca, &bonus)

        inimigos[i][0], inimigos[i][1] = forca, bonus
    }

    inimigos = mergesort(inimigos)

    for i, _ := range inimigos {
        if forcaKirito > inimigos[i][0] {
            forcaKirito += inimigos[i][1]
        } else {
            fmt.Println("NO")
            return
        }
    }

    fmt.Println("YES")
}
