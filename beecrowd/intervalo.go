package main

import "fmt"

func main() {
    var valor float64

    fmt.Scan(&valor)

    switch {
    case 0.0 <= valor && valor <= 25.0:
        fmt.Println("Intervalo [0,25]")
    case 25.0 < valor && valor <= 50.0:
        fmt.Println("Intervalo (25,50]")
    case 50.0 < valor && valor <= 75.0:
        fmt.Println("Intervalo (50,75]")
    case 75.0 < valor && valor <= 100.0:
        fmt.Println("Intervalo (75,100]")
    default:
        fmt.Println("Fora de intervalo")
    }
}
