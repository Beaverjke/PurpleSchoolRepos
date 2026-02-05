package main

import (
	"errors"
	"fmt"
)

const usdToEur = 0.8576 // стоимость Доллара в Евро
const usdToRub = 80.59  // стоимость Доллара в Рублях
const eurToRub = 93.84  // стоимость Евро в Рублях
const eurToUsd = 1.17   //
const rubToEur = 0.0106
const rubToUsd = 0.012

func getUserInput() (float64, float64, float64) {
	var fValue, sValue, tValue float64
	fmt.Println("Введи исходную валюту:\n1 - Евро\n2 - Доллар\n3 - Рубль")
	fmt.Scan(&fValue)
	fmt.Println("Введи сумму:")
	fmt.Scan(&sValue)
	fmt.Println("Введи валюту, в которую хочешь сконвертировать:")
	switch {
	case fValue == 1:
		fmt.Println("1 - Доллар\n2 - Рубль")
	case fValue == 2:
		fmt.Println("1 - Евро\n2 - Рубль")
	case fValue == 3:
		fmt.Println("1 - Евро\n2 - Доллар")
	}
	fmt.Scan(&tValue)
	return fValue, sValue, tValue
}
func checkRepeatCalc() bool {
	var uChoice int
	fmt.Println("Хочешь повторить вычисление?\n1 - Да\n0 - Нет")
	fmt.Scan(&uChoice)
	if uChoice == 1 {
		return true
	}
	return false
}

func calcValue(fValue, sValue, tValue float64) (float64, error) {
	var value float64

	switch {
	case fValue == 1 && tValue == 1:
		value = eurToUsd * sValue
	case fValue == 1 && tValue == 2:
		value = eurToRub * sValue
	case fValue == 2 && tValue == 1:
		value = usdToEur * sValue
	case fValue == 2 && tValue == 2:
		value = usdToRub * sValue
	case fValue == 3 && tValue == 1:
		value = rubToEur * sValue
	case fValue == 3 && tValue == 2:
		value = rubToUsd * sValue
	}
	if fValue == 0 || tValue == 0 {
		return 0, errors.New("такой валюты нет")
	} else if fValue > 3 && tValue < 1 || tValue > 3 && fValue < 1 {
		return 0, errors.New("ты опечатался, введи значения заново")
	}
	return value, nil
}

func main() {
	for {
		var startCurrency, sum, finishCurrency = getUserInput()
		var value, err = calcValue(startCurrency, sum, finishCurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Результат твоей конвертации: %.2f\n", value)

		isRepeat := checkRepeatCalc()
		if isRepeat != true {
			break
		}
	}
}
