package sortFuctions

import (
	"fmt"
	"os"
)

func BubbleSort() {
	var name string
	var a float64
	fmt.Print("Введите своё имя: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Println("Привет, " + name + "! Это реализация алгоритма пузырьковой сортировки на Go. Автор: Сёмин Александр.")

	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("Вы ввели число:", a)

	fmt.Println("Программа завершена, нажмите любую клавишу для выхода")
	fmt.Scanln()
}
