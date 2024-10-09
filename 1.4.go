package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worker читает данные из канала и выводит их в stdout, пока не получит сигнал завершения контекста
func worker(ctx context.Context, id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Получаем сигнал завершения
			fmt.Printf("Worker %d завершает работу.\n", id)
			return
		case data, ok := <-dataCh:
			if !ok {
				fmt.Printf("Worker %d завершает работу, канал закрыт.\n", id)
				return
			}
			fmt.Printf("Worker %d получил данные: %d\n", id, data)
		}
	}
}

func main() {
	fmt.Print("Укажите количество воркеров: ")
	var nw int
	fmt.Scanln(&nw)

	if nw <= 0 {
		fmt.Println("Некорректное количество воркеров. Укажите положительное число.")
		return
	}

	// Создаем канал для передачи данных
	dataCh := make(chan int)

	// Контекст для управления временем жизни воркеров
	ctx, cancel := context.WithCancel(context.Background())

	// WaitGroup для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Запуск N воркеров
	for i := 1; i <= nw; i++ {
		wg.Add(1)
		go worker(ctx, i, dataCh, &wg)
	}

	// Генерация данных в отдельной горутине
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dataCh <- rand.Intn(100)           // Отправляем случайные данные в канал
				time.Sleep(500 * time.Millisecond) // Имитируем задержку
			}
		}
	}()

	// Канал для перехвата сигнала завершения (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-sigCh
	fmt.Println("Получен сигнал завершения. Останавливаем воркеров...")

	// Отправляем сигнал на завершение контекста
	cancel()

	// Закрываем канал для данных
	close(dataCh)

	// Ожидаем завершения всех воркеров
	wg.Wait()

	fmt.Println("Все воркеры завершили работу.")
}
