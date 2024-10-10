package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6} // массив чисел

	inputCh := make(chan int)  // создаем канал для входных чисел
	outputCh := make(chan int) // создаем канал для выходных чисел

	// горутина для записи чисел в входной канал
	go func() {
		for _, el := range numbers {
			inputCh <- el // отправляем число в входной канал
		}
		close(inputCh) // закрываем входной канал после записи всех чисел
	}()

	// горутина для обработки чисел: умножаем на 2 и отправляем в выходной канал
	go func() {
		for num := range inputCh {
			outputCh <- num * 2 // умножаем число на 2 и отправляем в выходной канал
		}
		close(outputCh) // закрываем выходной канал после обработки всех чисел
	}()

	// читаем и выводим числа из выходного канала
	for n := range outputCh {
		fmt.Println(n) // выводим число в stdout
	}
}
