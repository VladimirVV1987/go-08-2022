/*
Урок 3. Операторы и управляющие конструкции, базовые функции
1. Доработать калькулятор из методички: больше операций и валидации данных 
(проверка деления на 0, возведение в степень, факториал) + форматирование строки при выводе дробного 
числа ( 2 знака после точки)

*/

package main

import (
	"fmt"
	"math"
	"os"
)

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var a, b, res float64
	var op string
	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)
	//a = 4
	fmt.Println("Введите второе число: ")
	fmt.Scanln(&b)
	//b = 3
	fmt.Println("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	//op = "!"

	switch op {
	case "+":
		res = a + b
		fmt.Println("Результат сложения: ", res)
	case "-":
		res = a - b
		fmt.Println("Результат вычитания: ", res)
	case "*":
		res = a * b
		fmt.Println("Результат умножения: ", res)
	case "/":
		//проверка деления на 0
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
		}
		res = a / b
		fmt.Println("Результат деления: ", res)
	case "^":
		res = math.Pow(a, b)
		fmt.Println("Результат возведения в степень: ", res)
	case "!":
		var aa uint = uint(a) //из float64 в uint
		var bb uint = uint(b) //из float64 в uint
		res2 := factorial(aa)
		fmt.Println("Результат операции факториал: ", aa, " равен ", res2)
		res22 := factorial(bb)
		fmt.Println("Результат операции факториал: ", bb, " равен ", res22)
		break
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}



}
