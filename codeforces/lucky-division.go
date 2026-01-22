package main

import "fmt"


func ehDivisivelPorSortudo(numero int) bool {
    numerosSortudos := []int{4, 7, 44, 47, 74, 77, 444, 447, 474, 477, 744, 747, 774, 777}

    for _, v := range numerosSortudos {
        if numero % v == 0 {
            return true
        }
    }

    return false
}

func main() {
    var numero int

    fmt.Scan(&numero)

    if ehDivisivelPorSortudo(numero) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
