package main

import "fmt"

func main() {
    var casos int

    fmt.Scan(&casos)

    for i := 0; i < casos; i++ {
        var n, k int

        fmt.Scan(&n, &k)

        fmt.Println(k + ((k - 1) / (n - 1)))
    }
}
