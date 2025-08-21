package main

import (
	"fmt"
	"math"
)

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Println("Введи свой рост в СМ: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введи свой вес в КГ: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}

func calcIMT(uWeight float64, uHeight float64) float64 {
	const IMTPower = 2
	IMT := uWeight / math.Pow(uHeight/100, 2)
	return IMT
}

func IMTOutput(imt float64) {
	result := fmt.Sprintf("Твой IMT: %.1f", imt)
	fmt.Println(result)
}

func main() {
	fmt.Println("_-_-_ Это калькулятор IMT _-_-_")
	IMT := calcIMT(getUserInput()) // userWeight, userHeight := getUserInput()
	if IMT < 16 {
		fmt.Println("Ты очень худой!")
	} else if IMT < 18.5 {
		fmt.Println("Ты худой!")
	} else if IMT < 25 {
		fmt.Println("Ты норм!")
	} else if IMT < 30 {
		fmt.Println("Ты ты пухлый!")
	} else {
		fmt.Println("Ты жирный!")
	}
	IMTOutput(IMT)
}
