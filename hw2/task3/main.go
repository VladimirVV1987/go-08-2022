/*
Урок 2
3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
*/

package main

import (
	"fmt"
)

func main() {
	var a, b, c, qwerty int
		
	fmt.Print("Введите трехзначное число:  ")
	fmt.Scanln(&qwerty)

	a = qwerty / 100 % 10
	b = qwerty / 10 % 10
	c = qwerty / 1 % 10
	
	fmt.Println("Сотни: ", a)
	fmt.Println("десятки: ", b)
	fmt.Println("единицы: ", c)

}
