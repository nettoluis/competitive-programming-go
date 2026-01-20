package main

import "fmt"

func main() {
    var qtd int

    fmt.Scan(&qtd)

    coordenadas := make(map[string]int)

    for i := 0; i < qtd; i++ {
        var xi, yi, zi int

        fmt.Scan(&xi, &yi, &zi)

        coordenadas["x"] += xi
        coordenadas["y"] += yi
        coordenadas["z"] += zi
    }

    if coordenadas["x"] + coordenadas["y"] + coordenadas["z"] == 0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
