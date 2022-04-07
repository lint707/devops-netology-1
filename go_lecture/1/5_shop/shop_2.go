package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	store := map[string]int{}
	prices := map[string]int{}
	statByCnt := map[string]int{}
	statByPrice := map[string]int{}
	operations := []string{}

	input := bufio.NewScanner(os.Stdin)

LOOP:
	for input.Scan() {
		text := input.Text()

		params := strings.Split(text, " ")

		cmd := params[0]
		params = params[1:]

		switch cmd {
		case "поставка":
			if len(params) != 2 {
				fmt.Println("надо 3 аргумента")
				continue LOOP
			}
			cnt, err := strconv.Atoi(params[1])
			if err != nil {
				fmt.Println("укажите число")
				continue LOOP
			}
			itemName := params[0]
			store[itemName] = store[itemName] + cnt
			operations = append(operations, text)
		case "цена":
			paramsStr := strings.Join(params, " ")
			var (
				itemName string
				price    int
			)
			_, err := fmt.Sscanf(paramsStr, "%s %d", &itemName, &price)
			if err != nil {
				fmt.Println("ошибка ввода данных:", err)
				continue LOOP
			}

			if _, exist := store[itemName]; !exist {
				fmt.Println("нет такого товар:", itemName)
				continue LOOP
			}
			prices[itemName] = price
			operations = append(operations, text)
		case "склад":
			fmt.Println("Статистика по складку")
			for itemName, cnt := range store {
				price := prices[itemName]
				fmt.Println(itemName, " количество:", cnt, "цена:", price)
			}
		case "продажа":
			if len(params) != 2 {
				fmt.Println("надо 3 аргумента")
				continue LOOP
			}
			cnt, err := strconv.Atoi(params[1])
			if err != nil {
				fmt.Println("укажите число")
				continue LOOP
			}
			itemName := params[0]

			inStore, exist := store[itemName]
			if !exist {
				fmt.Println("нет такого товар:", itemName)
				continue LOOP
			}
			if inStore < cnt {
				fmt.Println("на складе нету столько товара:", itemName)
				continue LOOP
			}

			price, exist := prices[itemName]
			if !exist {
				fmt.Println("еще не установлена цена:", itemName)
				continue LOOP
			}

			statByCnt[itemName] += cnt
			statByPrice[itemName] += cnt * price

			store[itemName] = store[itemName] - cnt

			opText := fmt.Sprintf("%s на сумму %d по цене %d", text, cnt*price, price)
			operations = append(operations, opText)
		case "стат":
			for itemName, sum := range statByPrice {
				total := statByCnt[itemName]
				var avgPrice float32 = float32(sum / total)
				fmt.Printf("продано %d %s на сумму %d, средняя цена: %f\n", total, itemName, sum, avgPrice)
			}
		case "операции":
			for i, opText := range operations {
				fmt.Println(i, opText)
			}
		}

		//fmt.Printf("params: %#v\n", params)
	}

	fmt.Println(store)
}
