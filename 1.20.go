package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// разделяем строку на слова
	words := strings.Fields(s)

	// переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// объединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"         // исходная строка
	reversed := reverseWords(input) // перевёрнутая строка по словам
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевернутая строка:", reversed)
}
