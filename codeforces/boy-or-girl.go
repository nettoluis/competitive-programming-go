package main

import "fmt"

func countChars(s string) int {
    uniqueChars := make(map[rune]bool)

    for _, char := range s {
        uniqueChars[char] = true
    }

    return len(uniqueChars)
}


func main() {
    var username string

    fmt.Scan(&username)

    numOfUniqueChars := countChars(username)
    switch {
        case numOfUniqueChars % 2 == 0:
            fmt.Println("CHAT WITH HER!")
        case numOfUniqueChars % 2 != 0:
            fmt.Println("IGNORE HIM!")
    }
}
