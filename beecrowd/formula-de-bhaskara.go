package main

import (
    "fmt"
    "math"
)

func calculaDelta(a, b, c float64) float64 {
    return (b * b) - (4.0 * a * c)
}

func calculaRaizes(a, b, delta float64) (r1 float64, r2 float64){
    r1 = (-b + math.Pow(delta, 0.5)) / (2.0 * a)
    r2 = (-b - math.Pow(delta, 0.5)) / (2.0 * a)

    return
}

func main() {
    var a, b, c float64

    fmt.Scan(&a, &b, &c)

    delta := calculaDelta(a, b, c)

    if delta >= 0.0 && a > 0.0 {
        raiz1, raiz2 := calculaRaizes(a, b, delta)
        fmt.Printf("R1 = %.5f\nR2 = %.5f\n", raiz1, raiz2)
    } else {
        fmt.Println("Impossivel calcular")
    }
}
