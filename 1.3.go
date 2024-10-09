package main

import (
	"fmt"
	"os"
	"sync"
)

// Sq выводит квадрат числа в stdout и отправляет его в канал
func Sq(n int, num chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Сообщаем, что горутина завершила свою работу
	s := n * n
	fmt.Fprintln(os.Stdout, s)
	num <- s // Отправляем результат в канал
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	num := make(chan int)

	// Проходим по массиву и для каждого числа вычисляем квадрат в отдельной горутине
	for _, el := range arr {
		wg.Add(1) // Увеличиваем счетчик активных горутин
		go Sq(el, num, &wg)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()  // Ожидаем завершения всех горутин
		close(num) // Закрываем канал
	}()

	sum := 0
	// Итерируем по каналу, пока он не закроется
	for el := range num {
		sum += el
	}

	fmt.Fprintln(os.Stdout, "Сумма квадратов:", sum)
}
