package main

import (
	"fmt"
	"time"
)

// MySleep реализует задержку на заданное количество времени.
func MySleep(duration time.Duration) {
	// используем time.After для создания канала, который будет сигнализировать о завершении
	<-time.After(duration)
}

func main() {
	fmt.Println("Начинаем ожидание...")
	MySleep(3 * time.Second) // ждём 3 секунды
	fmt.Println("Ожидание завершено!")
}
