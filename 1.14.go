package main

import (
	"fmt"
	"reflect"
)

func determineType(x interface{}) {
	// Проверка, является ли переменная каналом через пакет reflect
	if reflect.TypeOf(x).Kind() == reflect.Chan {
		fmt.Println("Тип: channel")
		return
	}

	// Используем type switch для остальных типов
	switch v := x.(type) {
	case int:
		fmt.Println("Тип: int, Значение:", v)
	case string:
		fmt.Println("Тип: string, Значение:", v)
	case bool:
		fmt.Println("Тип: bool, Значение:", v)
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	// Пример с переменной типа int
	determineType(42)
	// Пример с переменной типа string
	determineType("Hello, world!")
	// Пример с переменной типа bool
	determineType(true)
	// Пример с переменной типа канал
	ch := make(chan struct{})
	determineType(ch)
	// Пример с каналом другого типа
	ch2 := make(chan int)
	determineType(ch2)
	// Пример с переменной неизвестного типа
	determineType(3.14)
}
