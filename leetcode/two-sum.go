package main

func twoSum(nums []int, target int) []int {
    indice := make(map[int]int)

    for i, n := range nums {
        diferenca := target - n

        if idx, ok := indice[diferenca]; ok {
            return []int{idx, i}
        } else {
            indice[n] = i
        }  
    }

    return []int{-1, -1}
}
