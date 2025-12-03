package main

import "fmt"

func main() {
    var segundos int

    fmt.Scan(&segundos)

    restoHoras := segundos % 3600
    restoMinutos := restoHoras % 60
    fmt.Printf("%d:%d:%d\n", segundos / 3600, restoHoras / 60, restoMinutos)
}
