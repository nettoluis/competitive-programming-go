package main

import "fmt"

const Pi = 3.14159

func calculaArea(a float64) (valor float64) {
    valor = Pi * a * a
    return
}

func main() {
    var a float64

    fmt.Scanf("%f", &a)

    area := calculaArea(a)

    fmt.Printf("A=%.4f\n", area)
}
