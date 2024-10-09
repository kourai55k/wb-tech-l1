package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 1. Использование канала
	fmt.Println("1. Использование канала")

	// Канал для остановки
	stopCh := make(chan bool)

	// Увеличиваем счетчик WaitGroup
	wg.Add(1)

	// Запускаем горутину
	go func() {
		// Уменьшнаем счетчик
		defer wg.Done()
		for {
			select {
			// В канал поступил для остановки поступили данные
			case <-stopCh:
				fmt.Println("Остановка горутины")
				// Завершаем работу горутины
				return
			default:
				// Имитируем работу
				fmt.Println("Горутина работает")
				time.Sleep(400 * time.Millisecond)
			}
		}
	}()

	// Даём поработать горутине 2 секунды
	time.Sleep(2 * time.Second)

	// Отправляем в канал сигнал об остановке
	stopCh <- true

	// Ждём завершения горутин
	wg.Wait()

	fmt.Println("1. Горутина остановлена c помощью канала")
	fmt.Println()

	// 2. Использование контекста
	fmt.Println("2. Использование контекста")
	// Создаем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())

	// Увеличиваем счетчик WaitGroup
	wg.Add(1)

	// Запускаем горутину и передаём в неё контекст
	go func(ctx context.Context) {
		// Уменьшаем счетчик
		defer wg.Done()
		for {
			select {
			// Когда контекст отмениться, этот случай обработается и мы остановим горутиу
			case <-ctx.Done():
				fmt.Println("Остановка горутины")
				// Завершаем работу горутины
				return
			default:
				// Имитируем работу
				fmt.Println("Горутина работает")
				time.Sleep(400 * time.Millisecond)
			}
		}
	}(ctx)

	// Даём поработать 2 секунды
	time.Sleep(2 * time.Second)
	// Вызываем функцию отмены контекста
	cancel()
	// Ждём завершения горутин
	wg.Wait()

	fmt.Println("2. Горутина остановлена c помощью контекста")
	fmt.Println()

	// 3. Использование флага
	fmt.Println("3. Использование фалага")

	// Устанавливаем флаг
	flag := true

	// Увеличиваем счетчик WaitGroup
	wg.Add(1)
	// Запускаем горутину
	go func() {
		// Уменьшаем счетчик
		defer wg.Done()
		for {
			// Если флаг становится false, завершаем горутину
			if !flag {
				fmt.Println("Остановка горутины")
				// Завершаем горутину
				return
			}
			// Имитируем работу
			fmt.Println("Горутина работает")
			time.Sleep(400 * time.Millisecond)
		}
	}()

	// Даём поработать 2 секунды
	time.Sleep(2 * time.Second)
	// Устанавливаем флаг false
	flag = false
	// Ждём завершения горутин
	wg.Wait()
	fmt.Println("3. Горутина остановлена с помощью флага")
	fmt.Println()

	// 4. Исполбьзование паники(плохая практика)
	fmt.Println("4. Использование паники")

	wg.Add(1)
	// Запускаем горутину
	go func() {
		defer wg.Done()

		defer func() {
			if r := recover(); r != nil { // Обработка паники
				fmt.Println("Горутина остановлена из-за паники:", r)
			}
		}()

		for {
			fmt.Println("Горутина работает")
			time.Sleep(400 * time.Millisecond) // Имитируем работу

			// Эмулируем условие, когда нужно вызвать панику
			if time.Now().Second()%5 == 0 { // Каждые 5 секунд
				fmt.Printf("Сейчас %v часов %v минут и %v секунд\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
				panic("непредвиденная ошибка") // Вызываем панику
			}
		}
	}()

	// Даем горутине поработать 10 секунд
	//time.Sleep(10 * time.Second)
	wg.Wait()
	fmt.Println("4. Горутина остановлена с помощью паники")
}
