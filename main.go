package main

import (
	"./calculator"
	"./quadraticEquation"
	"./sortFuctions"
	"fmt"
	"os"
)

func main() {
	var programType int

	fmt.Println("Привет! Это набор программ. Автор: Сёмин Александр.")
	fmt.Println("Выберите программу:")
	fmt.Println("1 - Калькулятор")
	fmt.Println("2 - Квадратное уравнение")
	fmt.Println("3 - Сортировка пузырьком")
	fmt.Fscan(os.Stdin, &programType)

	switch programType {
	case 1:
		calculator.Calculator()
	case 2:
		quadraticEquation.QuadraticEquation()
	case 3:
		sortFuctions.BubbleSort()
	default:
		fmt.Println("Нет такой программы")
	}
}
