package main

import "fmt"

func main() {
    var primeira, segunda string

    fmt.Scan(&primeira, &segunda)

    if len(primeira) != len(segunda) {
        fmt.Println("NO")
        return
    }

    for i := 0; i < len(primeira); i++ {
        if primeira[i] != segunda[len(segunda) - i - 1] {
            fmt.Println("NO")
            return
        }
    }

    fmt.Println("YES")
}
