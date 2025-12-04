package main

import "fmt"

func main() {
    var dias int

    fmt.Scan(&dias)

    restoAnos := dias % 365
    restoMeses := restoAnos % 30

    fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", dias / 365, restoAnos / 30, restoMeses)
}
