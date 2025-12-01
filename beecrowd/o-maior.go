package main

import "fmt"

func encontraMaior(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func main() {
    var a, b, c int

    fmt.Scan(&a, &b, &c)

    maior := encontraMaior(encontraMaior(a, b), c)

    fmt.Printf("%d eh o maior\n", maior)
}
