package main

import "fmt"

func swap(s []int, i1, i2 int) {
    temp := s[i1]
    s[i1] = s[i2]
    s[i2] = temp
}

func bubbleSort(s []int) {
    swapped := true

    for swapped {
        swapped = false

        for i := 0; i < (len(s) - 1); i++ {
            if s[i] > s[i + 1] {
                swap(s, i, i + 1)
                swapped = true
            }
        }
    }
}

func main() {
    var a, b, c int

    fmt.Scan(&a, &b, &c)

    numeros := make([]int, 3)

    numeros[0], numeros[1], numeros[2] = a, b, c

    bubbleSort(numeros)

    for _, v := range numeros {
        fmt.Println(v)
    }

    fmt.Printf("\n%d\n%d\n%d\n", a, b, c)
}
