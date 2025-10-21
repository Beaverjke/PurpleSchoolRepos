package main

import (
	"fmt"
	"unicode"
)

func reverseArray(arr []int) []int {
	reversed := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}

	return reversed
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Print(reverseArray(arr1))

	str := "hhhhh20&&&&nd"

	digits := []rune{}

	for _, char := range str {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	fmt.Println(string(digits)) // 123456
}

//Массивы и слайс
//transactions := []int{1, 2, 3, 4} // Массив из чисел
//transactions = append(transactions, 5)
//newTrans := append(transactions, 15)
//fmt.Println(transactions)
//fmt.Println(newTrans)

//transactionsPartial := transactions[1:5]
//transactionsNewPartial := transactionsPartial[:1]
//banks := [3]string{"Sber", "T-Bank", "VTB"} // Массив из строк

//banks[0] = "Alpha" // Изменение значения части массива по индексу
//fmt.Println(banks)
//fmt.Println(transactions[1]) // Вывод ВТОРОГО! значения из массива

//partial := transactions[1:4] // Slice от 1(2) значения по индексу до последнего 4(5)
//fmt.Println(partial)

//transactionsNewPartial[0] = 10

//transactionsNewPartial = transactionsNewPartial[0:4]

//fmt.Println(transactions)
//fmt.Println(transactionsPartial)
//fmt.Println(transactionsNewPartial)
//fmt.Println(len(transactionsPartial), cap(transactionsPartial))
//fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial))

//tr1 := []int{1, 2, 3, 4}
//tr2 := []int{5, 6, 7, 8, 9}
//tr3 := append(tr1, tr2...)
//for index, value := range tr3 {
//	fmt.Println(index, value)
//}
