package main

import "fmt"

const Pi = 3.14159

func main() {
    var a, b, c float64

    fmt.Scan(&a, &b, &c)

    triangulo := (a * c) / 2.0
    circulo := Pi * c * c
    trapezio := (a + b) * c / 2
    quadrado := b * b
    retangulo := a * b

    fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", triangulo, circulo, trapezio, quadrado, retangulo)
}
