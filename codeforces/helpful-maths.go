package main

import (
    "fmt"
    "strings"
    "strconv"
)

func intCasting(s []string) []int {
    saida := make([]int, len(s))

    for i := 0; i < len(s); i++ {
        elemento, err := strconv.Atoi(s[i])
        if err != nil {
            panic(err)
        }

        saida[i] = elemento
    } 

    return saida
}

func swap(s []int, i1, i2 int) {
    temp := s[i1]
    s[i1] = s[i2]
    s[i2] = temp
} 

func bubbleSort(s []int) {
    for true {
        swapped := false
        for i := 0; i < (len(s) - 1); i++ {
            if s[i] > s[i + 1] {
                swap(s, i, i + 1)
                swapped = true
            }
        }

        if !swapped {
            break
        }
    }
}

func main() {
    var entrada string

    fmt.Scan(&entrada)
    nums := strings.Split(entrada,"+")

    numsInt := intCasting(nums)

    bubbleSort(numsInt)

    saida := ""
    for i := 0; i < len(nums); i++ {
        if i != 0 {
            saida += fmt.Sprintf("+%d", numsInt[i])
        } else {
            saida += fmt.Sprintf("%d", numsInt[i])
        }
    }

    fmt.Println(saida)
}
