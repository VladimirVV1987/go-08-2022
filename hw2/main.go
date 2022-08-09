//Урок 2.
//1. Напишите программу для вычисления площади прямоугольника.
//Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.

package main

import (
	"fmt"
)

func main() {
	var a, b, square float64
	var result string

	fmt.Println("-=Вычисление площади прямоугольника=-")

	fmt.Print("Введите длину прямоугольника:  ")
	fmt.Scanln(&a)

	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scanln(&b)

	square = a * b
	result = fmt.Sprintf("%.2f", square)
	fmt.Printf("Площадь прямоугольника: %s\n", result)
}
