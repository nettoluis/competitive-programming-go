package main

import "fmt"

const (
    PesoA = 3.5
    PesoB = 7.5
)

func calculaMedia(a, b float64) float64 {
    return ((a * PesoA) + (b * PesoB)) / (PesoA + PesoB)
}

func main() {
    var a, b float64

    fmt.Scan(&a, &b)

    media := calculaMedia(a, b)

    fmt.Printf("MEDIA = %.5f\n", media)
}
