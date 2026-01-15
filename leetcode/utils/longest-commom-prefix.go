package utils

func LongestCommonPrefix(strs []string) string {
    resultado := ""

    for i, _ := range strs[0] {
        for _, s := range strs {
            if i == len(s) || s[i] != strs[0][i] {
                return resultado
            }
        }

        resultado += string(strs[0][i])
    }

    return resultado
}
