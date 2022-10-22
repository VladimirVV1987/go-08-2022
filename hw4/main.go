/*
Урок 4. Сложные типы данных: массивы, слайсы и мапы
1. Познакомьтесь с алгоритмом сортировки вставками. Напишите приложение, которое принимает на вход набор целых чисел и
отдаёт на выходе его же в отсортированном виде.
Сохраните код, он понадобится нам в дальнейших уроках.
*/
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f int
	fmt.Println("Введите число a: ")
	fmt.Scanln(&a)
	fmt.Println("Введите число b: ")
	fmt.Scanln(&b)
	fmt.Println("Введите число c: ")
	fmt.Scanln(&c)
	fmt.Println("Введите число d: ")
	fmt.Scanln(&d)
	fmt.Println("Введите число e: ")
	fmt.Scanln(&e)
	fmt.Println("Введите число f: ")
	fmt.Scanln(&f)
	fmt.Println("Вы ввели следующие значения: ", a, b, c, d, e, f)

	/*
	   a := 2
	   b := 7
	   c := 5
	   d := 11
	   e := -1
	   f := 0

	*/
	arr := []int{a, b, c, d, e, f}
	InsertionSort(arr)
	fmt.Println(arr)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}
