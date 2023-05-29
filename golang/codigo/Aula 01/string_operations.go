package stringutil

import "strings"

// Função para inverter uma string
func ReverseString(input string) string {
    // Converter a string em um slice de caracteres
    chars := []rune(input)
    length := len(chars)

    // Inverter o slice de caracteres
    for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
        chars[i], chars[j] = chars[j], chars[i]
    }

    // Converter o slice de caracteres de volta para uma string
    return string(chars)
}

// Função para verificar se uma string é um palíndromo
func IsPalindrome(input string) bool {
    // Remover espaços em branco e converter para letras minúsculas
    cleanedInput := strings.ToLower(strings.ReplaceAll(input, " ", ""))

    // Comparar a string original com a versão invertida
    return cleanedInput == ReverseString(cleanedInput)
}
