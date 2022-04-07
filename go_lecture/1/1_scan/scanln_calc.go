package main

import (
	"fmt"
)

var (
	UserName = "v.romanov"
)

func main() {
	var totalOperations int
LOOP:
	for {
		// часть 1 - получение входных данных
		var (
			v1, v2 int
			op     string
		)
		fmt.Printf("число оператор число: ")
		// читаем значения в переменные
		_, err := fmt.Scanf("%d %s %d\n", &v1, &op, &v2)
		if err != nil {
			fmt.Println("ошибка ввода:", err)
			continue
		}

		// часть 2 - операции над данными
		// выбор операции на основе условия
		var result int
		switch op {
		case "-":
			result = v1 - v2
		case "*":
			result = v1 * v2
		case "END":
			fmt.Println("вы запросили завершение цикла")
			// завершаем выполнение цикла
			// LOOP для указаним что завершаем цикл, а не блок switch-case
			break LOOP
			fmt.Println("не выведется без LOOP")
		default:
			fmt.Println(op, "неизвестный оператор")
			continue
		}

		// часть 3 - вывод результата
		fmt.Println("операция", totalOperations+1)
		fmt.Println(v1, op, v2, "=", result)

		totalOperations++
	}
	// часть 3 - вывод результата
	fmt.Println("операций совершено:", totalOperations)
	fmt.Println("конец программы")
}
