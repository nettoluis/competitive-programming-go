package main

import "fmt"

func main() {
    var qtd int
    var partidas string

    fmt.Scan(&qtd, &partidas)

    anton, danik := 0, 0
    for i := 0; i < qtd; i++ {
        if string(partidas[i]) == "A" {
            anton++
        } else {
            danik++
        }
    }

    switch {
        case anton == danik:
            fmt.Println("Friendship")
        case anton > danik:
            fmt.Println("Anton")
        case anton < danik:
            fmt.Println("Danik")
    }


}
