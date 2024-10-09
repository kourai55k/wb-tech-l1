package main

import (
	"fmt"
	"math/big"
)

func main() {
	// задаем значения a и b
	a := int64(1 << 22) //  2^22
	b := int64(1 << 21) //  2^21

	// выполняем арифметические операции и выводим результаты сразу
	fmt.Printf("int64:\n")
	sum := a + b // сумма
	fmt.Printf("Сумма: %v\n", sum)
	sub := a - b // разность
	fmt.Printf("Разность: %v\n", sub)
	mul := a * b // произведение
	fmt.Printf("Произведение: %v\n", mul)
	// частное
	if b != 0 {
		div := a / b
		fmt.Printf("Частное: %v\n", div)
	} else {
		fmt.Println("Ошибка: Деление на ноль")
	}

	// int64 может хранить числа размером до 2^63,
	// если нам надо хранить числа больше, можно использовать math/big
	// у big.Int нет явного ограничения, оно зависит от количества доступной
	// оперативной памяти устройства

	// создаём переменные a и b с числами больше 2^20
	a1 := new(big.Int)
	b1 := new(big.Int)

	// задаем значение 1 для a1 и b1 с использованием SetInt64
	a1.SetInt64(1) // устанавливаем значение 1
	b1.SetInt64(1) // устанавливаем значение 1
	a1.Lsh(a1, 22) // 1 << 22 (это 2^22)
	b1.Lsh(b1, 21) // 1 << 21 (это 2^21)

	fmt.Printf("big.Int:\n")
	// сумма
	sum1 := new(big.Int).Add(a1, b1)
	fmt.Printf("Сумма: %v\n", sum1)
	// разность
	sub1 := new(big.Int).Sub(a1, b1)
	fmt.Printf("Разность: %v\n", sub1)
	// произведение
	mul1 := new(big.Int).Mul(a1, b1)
	fmt.Printf("Произведение: %v\n", mul1)
	// частное
	if b1.Cmp(big.NewInt(0)) != 0 { // проверяем, что b1 не равно нулю
		div1 := new(big.Int).Div(a1, b1)
		fmt.Printf("Частное: %v\n", div1)
	} else {
		fmt.Println("Ошибка: Деление на ноль")
	}
}
