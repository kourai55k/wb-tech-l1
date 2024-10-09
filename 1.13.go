package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	fmt.Println("До обмена:", a, b)

	// В Go есть встроенный синтаксис для обмена значениями
	a, b = b, a

	fmt.Println("После обмена:", a, b)
}
