/*
Урок 5. Сложные типы данных: структуры, функции и методы
2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.** (Оно так же должно быть рекурсивным!)**
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
	var mymap = make(map[uint]int)
	mymap[2] = 1
	mymap[3] = 2
	mymap[4] = 3
	mymap[5] = 5
	mymap[6] = 8
	mymap[7] = 13
	mymap[8] = 21

	var a uint
	fmt.Println("Введите число a: ")
	fmt.Scanln(&a)

	if mymap[a] > 8 {
		fmt.Println("Результат из мап: ", mymap[a])
	} else {
		fmt.Println("Результат из функции Фиббоначи: ", fibbonachi(a))
	}

}
