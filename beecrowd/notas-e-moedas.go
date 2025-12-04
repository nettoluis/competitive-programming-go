package main

import "fmt"

func main() {
    cedulas := []int{10000, 5000, 2000, 1000, 500, 200}
    moedas := []int{100, 50, 25, 10, 5, 1}
    var v float64

    fmt.Scan(&v)
    valor := int(v * 100)

    fmt.Println("NOTAS:")
    for _, cedula := range cedulas {
        fmt.Printf("%d nota(s) de R$ %d.00\n", valor / cedula, cedula / 100)
        valor = valor % cedula
    }
    fmt.Println("MOEDAS:")
    for i, moeda := range moedas {
        if i != 0 {
            fmt.Printf("%d moeda(s) de R$ 0.%02d\n", valor / moeda, moeda)
        } else {
            fmt.Printf("%d moeda(s) de R$ 1.00\n", valor / moeda)
        }
        valor = valor % moeda
    }
}
