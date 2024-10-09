package main

import (
	"fmt"
	"strings"
)

// isUnique проверяет, все ли символы в строке уникальны.
func isUnique(s string) bool {
	// создаем карту для хранения символов
	m := make(map[string]int)
	// приводим строку к нижнему регистру
	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok { // проверяем, встречался ли символ ранее
			// если да, значит символ не уникальный
			return false // возвращаем false
		}
		m[string(s[i])]++ // добавляем символ в карту
	}
	return true // если все символы уникальны, возвращаем true
}

func main() {
	s := "abcd"              // строка для проверки
	fmt.Println(isUnique(s)) // выводим результат проверки
}
