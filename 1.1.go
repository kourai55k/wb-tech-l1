package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

// Speak - метод структуры Human
func (h Human) Speak() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Action встраивает структуру Human, тем самым автоматически встраивает и её методы.
type Action struct {
	Human      // Встраивание структуры Human
	ActionType string
}

// DoAction - метод структуры Action
func (a Action) DoAction() {
	fmt.Printf("%s is doing action: %s\n", a.Name, a.ActionType)
}

func main() {
	// Создаем экземпляр структуры Action и устанавливаем поля
	a := Action{
		Human: Human{
			Name: "Robert",
			Age:  23,
		},
		ActionType: "programming",
	}

	// Вызов метода структуры Human
	a.Human.Speak()

	// Вызов метода структуры Human через Action
	a.Speak()

	// Вызов метода структуры Action
	a.DoAction()
}
