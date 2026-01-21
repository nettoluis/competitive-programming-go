package main

import "fmt"

func main() {
    var n, alturaMuro int
    
    fmt.Scan(&n, &alturaMuro)

    alturas := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&alturas[i])
    }

    largura := 0
    for _, v := range alturas {
        if v <= alturaMuro {
            largura++
        } else {
            largura += 2
        }
    }

    fmt.Println(largura)
}
