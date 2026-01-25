package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    resposta := max(
        a + b + c,
        (a + b) * c,
        a * (b + c),
        a * b * c,
    )

    fmt.Println(resposta)
}
