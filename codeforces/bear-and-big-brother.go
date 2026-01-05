package main

import (
    "fmt"
)

func main() {
    var limak, bob int64

    fmt.Scan(&limak, &bob)

    contador := 0
    for limak <= bob {
        limak *= 3
        bob *= 2
        contador++
    }

    fmt.Println(contador)
}
