package main

import "fmt"

func main() {
    cedulas := []int{100, 50, 20, 10, 5, 2, 1}
    var valor int

    fmt.Scan(&valor)

    fmt.Println(valor)
    for _, cedula := range cedulas {
        fmt.Printf("%d nota(s) de R$ %d,00\n", valor / cedula, cedula)
        valor = valor % cedula
    }
}
