package main

import "fmt"

func main() {

	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Начальный массив:", array)

	slice := array[1:5]
	fmt.Println("Начальный срез:", slice)

	array[1] = 22
	fmt.Println("Срез после изменения в массиве:", slice)

	slice[1] = 33

	fmt.Println("Массив после изменения в срезе:", array)
}
