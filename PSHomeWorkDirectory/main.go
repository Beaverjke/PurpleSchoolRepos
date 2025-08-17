package main

import (
	"fmt"
)

func getUserInput() (string, string, string) {
	var fValue string
	var sValue string
	var tValue string
	fmt.Print("Введи валюту: ")
	fmt.Scan(&fValue)
	fmt.Print("Введи валюту в которую производим конвертацию: ")
	fmt.Scan(&sValue)
	fmt.Print("Введи таргет: ")
	fmt.Scan(&tValue)
	return fValue, sValue, tValue
}

func calcValue(iValue float64, fCurrancy string, sCurrancy string) float64 {
	return 1
}
func main() {
	const usdToEur = 0.8541              // стоимость Доллара в Евро
	const usdToRub = 80.02               // стоимость Доллара в Рублях
	const eurToRub = usdToRub / usdToEur // стоимость Евро в Рублях

	// fmt.Printf("Евро равен: %.2f", eurToRub)

}
