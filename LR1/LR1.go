package main

import (
	"fmt"
	"time"
)

func main() {
	//Задание 1
	var now = time.Now()
	fmt.Println("Дата и время сейчас: ", now)

	//Задание 2
	var (
		i int     = 64
		f float64 = 6.4
		s string  = "строка"
		b bool
	)

	fmt.Println("Целое число: ", i)
	fmt.Println("Плавающая точка: ", f)
	fmt.Println("Строка: ", s)
	fmt.Println("Логическая: ", b)

	//Задание 3
	var r = 56
	fmt.Println("Вывод краткой формой: ", r)

	//Задание 4
	var j = 64
	var result = j + i
	fmt.Println("Сложение 2 целых чисел: ", result)

	//Задание 5
	var f2 = 4.6
	var result_fsum float64 = f2 + f
	var result_fmin float64 = f - f2
	fmt.Println("Сумма с плавающей точкой:", result_fsum)
	fmt.Printf("Разность с плавающей точкой: %.3f \n", result_fmin)

	var f3 = 4.7
	var result_avg = (f + f2 + f3) / 3
	fmt.Printf("Среднее значение 3 чисел: %.3f", result_avg)
}
