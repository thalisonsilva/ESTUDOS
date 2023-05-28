package main

import (
    "fmt"
    "stringutil"
)

func main() {
    input := "radar"

    reversed := stringutil.ReverseString(input)
    fmt.Println("String invertida:", reversed)

    isPalindrome := stringutil.IsPalindrome(input)
    fmt.Println("É palíndromo?", isPalindrome)
}
