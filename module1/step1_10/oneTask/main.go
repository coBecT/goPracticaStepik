package main

import (
	"fmt"
)

//Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.
//
//Sample Input:
//
//
//Sample Output:
//
//1
//4
//9
//16
//25
//36
//49
//64
//81
//100

func main() {
	for i, a := 1, 0; i <= 10; i++ {
		a = i * i
		fmt.Println(a)

	}
}
