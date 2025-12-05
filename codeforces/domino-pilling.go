package main

import "fmt"

func main() {
    var linhas, colunas int

    fmt.Scan(&linhas, &colunas)

    fmt.Printf("%d\n", (linhas * colunas) / 2)
}
