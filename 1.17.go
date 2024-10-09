package main

import "fmt"

func binarySearch(nums []int, t int) int {
	left, right := 0, len(nums)-1 // инициализируем левую и правую границу поиска

	for left <= right { // пока левая граница меньше или равна правой
		mid := left + (right-left)/2 // вычисляем середину массива

		if nums[mid] == t { // если элемент в середине равен искомому
			return mid // возвращаем индекс найденного элемента
		}

		if nums[mid] < t { // если элемент меньше искомого
			left = mid + 1 // ищем в правой части массива
		} else { // если элемент больше искомого
			right = mid - 1 // ищем в левой части массива
		}
	}
	return -1 // возвращаем -1, если элемент не найден
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден.\n", target)
	}
}
