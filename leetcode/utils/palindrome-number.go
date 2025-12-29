package utils

import "fmt"

func IsPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    xCopia := fmt.Sprintf("%d", x)
    y := ""

    for i := (len(xCopia) - 1); i >= 0; i-- {
        y += string(xCopia[i])
    }

    return y == xCopia
}
