package main

import "fmt"

func main() {
	transactions := [5]int{1, 2, 3, 4, 5}       // Массив из чисел
	banks := [3]string{"Sber", "T-Bank", "VTB"} // Массив из строк

	banks[0] = "Alpha" // Изменение значения части массива по индексу
	fmt.Println(banks)
	fmt.Println(transactions[1]) // Вывод ВТОРОГО! значения из массива

	partial := transactions[1:] // Slice от 1(2) значения по индексу до последнего 4(5)
	fmt.Println(partial)

}
