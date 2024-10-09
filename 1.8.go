package main

import (
	"fmt"
)

// setBit устанавливает i-тый бит в 1(value=true) или 0(value=false)
func setBit(num int64, i int, value bool) int64 {
	if value {
		// Устанавливаем i-й бит в 1
		return num | (1 << i)
	} else {
		// Устанавливаем i-й бит в 0
		return num &^ (1 << i)
	}
}

func main() {
	var num int64 = 42 // Пример числа
	fmt.Printf("Исходное число: %d\n", num)

	i := 1                     // Индекс бита для изменения
	num = setBit(num, i, true) // Установить 1
	fmt.Printf("После установки %d-го бита в 1: %d\n", i, num)

	i = 3
	num = setBit(num, i, false) // Установить 0
	fmt.Printf("После установки %d-го бита в 0: %d\n", i, num)
}
