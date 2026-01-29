package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)

    anterior := 0
    contador := 0
    for i := 0; i < n; i++ {
        var arranjo int

        fmt.Scan(&arranjo)

        if arranjo != anterior {
            anterior = arranjo
            contador++
        }
    }

    fmt.Println(contador)
}
