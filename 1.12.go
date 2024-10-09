package main

import (
	"fmt"
)

type StringSet struct {
	elements map[string]struct{}
}

// Создание нового множества
func NewStringSet() *StringSet {
	return &StringSet{elements: make(map[string]struct{})}
}

// Добавление элемента в множество
func (s *StringSet) Add(element string) {
	s.elements[element] = struct{}{}
}

// Проверка наличия элемента в множестве
func (s *StringSet) Contains(element string) bool {
	_, exists := s.elements[element]
	return exists
}

// Вывод элементов множества
func (s *StringSet) Items() []string {
	items := []string{}
	for key := range s.elements {
		items = append(items, key)
	}
	return items
}

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество и добавляем элементы
	set := NewStringSet()
	for _, str := range strings {
		set.Add(str)
	}

	// Вывод уникальных элементов
	fmt.Println("Уникальные элементы:", set.Items())
}
