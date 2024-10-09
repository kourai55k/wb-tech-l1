package main

import (
	"fmt"
)

func reverseString(s string) string {
	// преобразуем строку в срез рун (для поддержки unicode)
	runes := []rune(s)

	// переворачиваем руну
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) // возвращаем перевёрнутую строку
}

func main() {
	input := "главрыба"              // исходная строка
	reversed := reverseString(input) // перевёрнутая строка
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевернутая строка:", reversed)
}
