/*
Урок 5. Сложные типы данных: структуры, функции и методы
1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
*/
package main

import (
	"fmt"
)

func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
func main() {
	var a uint
	fmt.Println("Введите число a: ")
	fmt.Scanln(&a)
	fmt.Println("Результат: ", fibbonachi(a))

}
