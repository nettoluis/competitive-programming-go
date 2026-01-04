package main

import "fmt"

func main() {
    var x, y float64

    fmt.Scan(&x, &y)

    switch {
        case x == 0 && y == 0:
            fmt.Println("Origem")
        case x == 0:
            fmt.Println("Eixo Y")
        case y == 0:
            fmt.Println("Eixo X")
        case x > 0:
            if y > 0 {
                fmt.Println("Q1")
            } else {
                fmt.Println("Q4")
            }
        case x < 0:
            if y > 0 {
                fmt.Println("Q2")
            } else {
                fmt.Println("Q3")
            }
    }
}
