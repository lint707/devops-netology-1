package main

import (
	"fmt"
)

var inputInt int

func main() {
	/*
		y = f(x)

		input = scanln()
	*/

	// общая сумма
	var sum int
	for {
		// часть 1 - получение входных данных
		fmt.Printf("введите число: ")
		// var input string
		// читаем значение в переменную
		n, err := fmt.Scanln(&inputInt)
		fmt.Println("вы ввели", inputInt, n, err)

		// часть 2 - операции над данными
		// до конца блока
		// if input == "END" {
		// 	// завершаем выполнение цикла
		// 	fmt.Println("вы запросили завершение цикла")
		// 	break
		// }

		// // преобразуем строку к числу
		// // var err error
		// inputInt, err = strconv.Atoi(input)
		// if err != nil {
		// 	fmt.Println(input, "это не число:", err)
		// 	continue
		// }

		fmt.Printf("%d + %d = %d\n", sum, inputInt, sum+inputInt)
		// добавлем введенное число к общей сумме
		sum = sum + inputInt
	}
	// часть 3 - вывод результата
	fmt.Println("итоговая сумма", sum, inputInt)
	fmt.Println("конец программы")
}
