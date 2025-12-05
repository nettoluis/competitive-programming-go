package main

import (
    "fmt"
    "math"
)

func main() {
    matriz := make([][]int, 5)
    
    for i := 0; i < 5; i++ {
        linha := make([]int, 5)
        for j := 0; j < 5; j++ {
            fmt.Scan(&linha[j])
        }
        matriz[i] = linha
    }

    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            if matriz[i][j] == 1 {
                fmt.Println(math.Abs(float64(2 - i)) + math.Abs(float64(2 - j)))
            }
        }
    }

}
