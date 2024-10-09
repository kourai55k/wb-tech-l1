package main

import (
	"fmt"
)

func main() {
	// исходный слайс
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный слайс:", slice)

	// индекс элемента, который нужно удалить (30)
	i := 4

	// удаляем i-й элемент из слайса
	slice = removeElement(slice, i)

	// выводим слайс после удаления
	fmt.Println("Слайс после удаления:", slice)
}

// removeElement - функция для удаления элемента по индексу i
func removeElement(slice []int, i int) []int {
	// проверка на допустимость индекса
	if i < 0 || i >= len(slice) {
		fmt.Println("Индекс недопустим")
		return slice // возвращаем исходный слайс, если индекс недопустим
	}

	// удаляем элемент, объединяя две части слайса
	return append(slice[:i], slice[i+1:]...) // добавляем элементы до и после удаляемого
}
