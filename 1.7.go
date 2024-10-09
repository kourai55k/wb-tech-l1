package main

import (
	"fmt"
	"sync"
)

// SafeMap - структура, которая обеспечивает безопасный доступ к карте
type SafeMap struct {
	mu   sync.Mutex     // Мьютекс для синхронизации доступа к карте
	data map[int]string // Внутренний map для хранения значений
}

// Set - метод для записи данных в мапу с защитой от гонок данных
func (s *SafeMap) Set(key int, value string) {
	s.mu.Lock()                                   // Блокируем мьютекс перед записью в карту
	defer s.mu.Unlock()                           // Освобождаем мьютекс после завершения записи
	fmt.Printf("key: %d value: %s\n", key, value) // Выводим записываемое значение
	s.data[key] = value                           // Записываем значение в карту
}

// Get - метод для чтения данных из мапы с защитой от гонок данных
func (s *SafeMap) Get(key int) string {
	s.mu.Lock()         // Блокируем мьютекс перед чтением из карты
	defer s.mu.Unlock() // Освобождаем мьютекс после завершения чтения
	return s.data[key]  // Возвращаем значение из карты
}

func main() {
	safeMap := SafeMap{data: make(map[int]string)} // Инициализируем SafeMap с пустым внутренним map
	var wg sync.WaitGroup                          // WaitGroup для ожидания завершения горутин

	// Запуск нескольких горутин для записи в карту
	for i := 0; i < 10; i++ {
		wg.Add(1) // Увеличиваем счетчик горутин
		go func(i int) {
			defer wg.Done()                               // Уменьшаем счетчик после завершения горутины
			safeMap.Set(i, fmt.Sprintf("Value %d", i*10)) // Записываем значение в SafeMap
		}(i) // Передаем значение i в горутину
	}

	wg.Wait() // Ждем завершения всех горутин

	// Вывод значений из карты
	for i := 0; i < 10; i++ {
		fmt.Println(safeMap.Get(i)) // Читаем и выводим значения из SafeMap
	}
}
