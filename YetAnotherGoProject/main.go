package main

import (
	"errors"
	"fmt"
)

func checkRestart() bool {
	var userAnswerValue int
	fmt.Println(`Ты хочешь добавить транзакции?
	1 - Продолжить
	0 - Стоп`)
	fmt.Scan(&userAnswerValue)

	switch {
	case userAnswerValue == 1:
		return true
	}
	return false
}

func getUserTransactions() (float64, error) {
	var userTransactions float64
	fmt.Println("Введи значение одной транзакции: ")
	fmt.Scan(&userTransactions)
	if userTransactions == 0 {
		return userTransactions, errors.New("не найдено транзакций")
	}
	return userTransactions, nil
}

func transSum(transactions []float64) float64 {
	var sum float64
	sum = 0.0
	for _, value := range transactions {
		sum += value
	}
	return sum
}

func main() {
	exmp := make([]string, 0, 3)
	exmp = append(exmp, "1")
	exmp = append(exmp, "2")
	exmp = append(exmp, "3")
	fmt.Println(exmp)

	transactionsMain := []float64{}
	var transactionsSum float64
	for {
		transactions, err := getUserTransactions()
		if err != nil {
			fmt.Println(err)
			break
		}
		transactionsMain = append(transactionsMain, transactions)

		isRepeatCalc := checkRestart()
		if !isRepeatCalc {
			break
		}
	}

	transactionsSum = transSum(transactionsMain)

	fmt.Printf("Твой Баланс: %.1f", transactionsSum)
}
