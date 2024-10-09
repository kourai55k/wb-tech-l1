package main

import (
	"fmt"
	"math"
)

// Point - структура точки с инкапсулированными полями для координат
type Point struct {
	x float64
	y float64
}

// New - конструктор для создания новой точки
func New(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// GetX - геттер для поля x
func (p *Point) GetX() float64 {
	return p.x
}

// GetY - геттер для поля y
func (p *Point) GetY() float64 {
	return p.y
}

// SetX - сеттер для поля x
func (p *Point) SetX(x float64) {
	p.x = x
}

// SetY - сеттер для поля y
func (p *Point) SetY(y float64) {
	p.y = y
}

// Distance - метод для вычисления расстояния между двумя точками
func (p *Point) Distance(other *Point) float64 {
	dx := p.GetX() - other.GetX()
	dy := p.GetY() - other.GetY()
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// создаем две точки
	point1 := New(3, 4)
	point2 := New(0, 0)

	point2.SetX(5)
	point2.SetY(6)

	// вычисляем расстояние между ними
	distance := point1.Distance(point2)

	// вывод результата
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
