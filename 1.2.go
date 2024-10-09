package main

//
//import (
//	"fmt"
//	"os"
//	"sync"
//)
//
//// Sq выводит квадрат числа в stdout
//func Sq(n int, wg *sync.WaitGroup) {
//	defer wg.Done() // Сообщаем, что горутина завершила свою работу
//	fmt.Fprintln(os.Stdout, n*n)
//}
//
//func main() {
//	arr := []int{2, 4, 6, 8, 10}
//	var wg sync.WaitGroup
//
//	// Проходим в цикле for по массиву и для каждого числа вычисляем квадрат в отдельной горутине
//	for _, el := range arr {
//		wg.Add(1) // Увеличиваем счетчик активных горутин
//		go Sq(el, &wg)
//	}
//
//	wg.Wait() // Ожидаем завершения всех горутин
//}
