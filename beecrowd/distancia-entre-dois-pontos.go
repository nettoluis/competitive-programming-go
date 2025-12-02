package main

import (
    "fmt"
    "math"
)

func distanciaEntreDoisPontos(x1, y1, x2, y2 float64) float64 {
    return math.Pow((math.Pow(x1 - x2, 2) + math.Pow(y1 - y2, 2)) , 0.5)
}

func main() {
    var x1, y1, x2, y2 float64

    fmt.Scan(&x1, &y1, &x2, &y2)

    saida := distanciaEntreDoisPontos(x1, y1, x2, y2)

    fmt.Printf("%.4f\n", saida)
}
