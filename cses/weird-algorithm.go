package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    saida := fmt.Sprintf("%d", n)
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
            saida += " " + fmt.Sprintf("%d", n)
        } else {
            n = (n * 3) + 1
            saida += " " + fmt.Sprintf("%d", n)
        }
    }

    fmt.Println(saida)
}
