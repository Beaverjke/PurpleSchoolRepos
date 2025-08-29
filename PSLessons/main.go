package main

import (
	"errors"
	"fmt"
	"math"
)

func calcIMT(uWeight float64, uHeight float64) (float64, error) {
	if uWeight <= 0 {
		return 0.0, errors.New("не указан вес")
	} else if uHeight <= 0 {
		return 0.0, errors.New("не указана высота")
	}
	const IMTPower = 2
	IMT := uWeight / math.Pow(uHeight/100, 2)
	return IMT, nil
}

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered", r)
	//	}
	//}()
	fmt.Println("_-_-_ Это калькулятор IMT _-_-_")
	for {
		IMT, err := calcIMT(getUserInput()) // userWeight, userHeight := getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
			//panic(err)
		}
		IMTOutput(IMT)

		isRepeatCalc := checkRestart()
		if !isRepeatCalc {
			break
		}
	}
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Println("Введи свой рост в СМ: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введи свой вес в КГ: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}

func IMTOutput(imt float64) {
	result := fmt.Sprintf("Твой IMT: %.1f", imt)
	fmt.Println(result)

	switch {
	case imt < 16:

		fmt.Println("Ты очень худой!")
	case imt < 18.5:

		fmt.Println("Ты худой!")
	case imt < 25:

		fmt.Println("Ты норм!")
	case imt < 30:
		fmt.Println("Ты пухлый!")
	default:
		fmt.Println("Ты жирный!")
	}
}

func checkRestart() bool {
	var userAnswerValue int
	fmt.Println(`Ты хочешь запустить калькулятор снова?
	1 - Продолжить
	0 - Стоп`)
	fmt.Scan(&userAnswerValue)
	switch {
	case userAnswerValue == 1:
		return true
	}
	return false

	//if userAnswerValue == 1 {
	//	return true
	//}
	//return false
}
