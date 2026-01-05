package main

import "fmt"

func main() {
    var posicao int64

    fmt.Scan(&posicao)

    if posicao % 5 == 0 {
        fmt.Println(posicao / 5)
    } else {
        fmt.Println((posicao / 5) + 1)
    }
}
