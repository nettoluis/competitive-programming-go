package main

import "fmt"

func main() {
    var distancia, litros float64

    fmt.Scan(&distancia, &litros)

    eficiencia := distancia / litros

    fmt.Printf("%.3f km/l\n", eficiencia)
}
