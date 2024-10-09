package main

import "fmt"

// quickSort сортирует массив целых чисел с использованием алгоритма быстрой сортировки.
func quickSort(nums []int) []int {
	// Нечего сортировать, если длина массива 1 или 0
	if len(nums) < 2 {
		return nums
	}

	// Выбираем опорный элемент, здесь выбираем элемент по середине массива
	pivot := nums[len(nums)/2]

	// Инициализация срезов для хранения элементов меньше и больше опорного
	left := []int{}
	right := []int{}

	// Проходим по каждому элементу массива
	for _, el := range nums {
		// Если элемент меньше опорного, добавляем его в левый срез
		if el < pivot {
			left = append(left, el)
		} else if el > pivot { // Если элемент больше опорного, добавляем его в правый срез
			right = append(right, el)
		}
	}

	// Рекурсивно сортируем левый и правый срезы и объединяем с опорным элементом
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	// Пример массива для сортировки
	arr := []int{34, 7, 23, 32, 5, 62}
	fmt.Println("Исходный массив:", arr)

	// Вызов функции quickSort для сортировки массива
	sortedArr := quickSort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
