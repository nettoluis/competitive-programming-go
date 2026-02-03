package main

import "fmt"

func main() {
    var qtd int

    fmt.Scan(&qtd)

    times := make(map[string]int)
    for i := 0; i < qtd; i++ {
        var time string

        fmt.Scan(&time)

        times[time]++
    }

    maior, timeVencedor := 0, ""
    for k := range times {
        if times[k] >= maior {
            maior = times[k]
            timeVencedor = k
        }
    } 

    fmt.Println(timeVencedor)
}
