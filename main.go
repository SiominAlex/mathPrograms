package main

import (
	"fmt"
	"mathPrograms/calculator"
	"mathPrograms/quadraticEquation"
	"os"
)

func main() {
	var programType int

	fmt.Println("Привет! Это набор программ. Автор: Сёмин Александр.")
	fmt.Println("Выберите программу:")
	fmt.Println("1 - Калькулятор")
	fmt.Println("2 - Квадратное уравнение")
	fmt.Fscan(os.Stdin, &programType)

	switch programType {
	case 1:
		calculator.Calculator()
	case 2:
		quadraticEquation.QuadraticEquation()
	default:
		fmt.Println("Нет такой программы")
	}
}
