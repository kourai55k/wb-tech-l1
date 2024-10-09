package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Print("Укажите время работы программы: ")
	var duration int
	fmt.Scanln(&duration)

	// Создаем контекст с тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel() // Отменяем контекст после завершения

	// канал для данных
	data := make(chan int)

	// Запуск горутины для записи в канал
	go func() {
		for i := 0; ; i++ {
			select {
			// отправляем данные в канал data
			case data <- i:
				time.Sleep(100 * time.Millisecond) // Симуляция работы
			// когда время выйдет, вызовется функция Done() и мы завершим отправку данных
			case <-ctx.Done():
				fmt.Println("Writer stopped.")
				return
			}
		}
	}()

	// Чтение данных из канала
	go func() {
		for {
			select {
			// принимаем данные из канала data
			case d := <-data:
				fmt.Printf("Received: %d\n", d)
			// когда время выйдет, вызовется функция Done() и мы завершим отправку данных
			case <-ctx.Done():
				fmt.Println("Reader stopped.")
				return
			}
		}
	}()

	// Ожидание завершения контекста
	<-ctx.Done()
	fmt.Println("Время вышло")
}
