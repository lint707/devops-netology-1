package main

import "fmt"

func main() {
	// цикл, выполняет блок кода несколько раз
	for {
		var input string
		// читаем значение в переменную
		fmt.Scanln(&input)
		fmt.Println("вы ввели", input)
		if input == "END" {
			// завершаем выполнение цикла
			fmt.Println("вы запросили завершение цикла")
			break
		}
	}
	fmt.Println("конец программы")
}
