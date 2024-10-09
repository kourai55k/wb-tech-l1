package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val int64 // значение счётчика, которое будет изменяться атомарно
}

func (c *Counter) Increment() {
	// атомарное добавление 1 к значению счётчика
	atomic.AddInt64(&c.val, 1)
}

func (c *Counter) Value() int64 {
	// атомарное чтение текущего значения счётчика
	return atomic.LoadInt64(&c.val)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{} // создаём экземпляр структуры-счётчика

	// запускаем 100 горутин, каждая из которых увеличивает значение счётчика
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment() // инкрементируем счётчик в каждой горутине
		}()
	}

	// ожидаем завершения всех горутин
	wg.Wait()

	// выводим итоговое значение счётчика
	fmt.Println("Итоговое значение счётчика:", counter.Value())
}
