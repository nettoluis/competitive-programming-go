package main

import "fmt"

func main() {
    var qtd, tempo int
    var entrada string

    fmt.Scan(&qtd, &tempo)

    fmt.Scan(&entrada)

    fila := []rune(entrada)

    for tempo > 0 {
        for i := 0; i < len(fila) - 1; i++ {
            if fila[i] == 'B' && fila[i+1] == 'G' {
                fila[i], fila[i+1] = fila[i+1], fila[i]
                i++
            }
        } 

        tempo--
    }

    fmt.Println(string(fila))
}
