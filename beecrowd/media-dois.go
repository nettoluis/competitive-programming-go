package main

import "fmt"

const (
    PesoA = 2
    PesoB = 3
    PesoC = 5
) 
func calculaMedia(a, b, c float64) float64 {
    return ((PesoA * a) + (PesoB * b) + (PesoC * c)) / (PesoA + PesoB + PesoC)
}

func main() {
    var a, b, c float64

    fmt.Scan(&a, &b, &c)

    media := calculaMedia(a, b, c)

    fmt.Printf("MEDIA = %.1f\n", media)
}
