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
		case "склад":
			fmt.Println("Статистика по складку")
			for itemName, cnt := range store {
				fmt.Println(itemName, " количество:", cnt)
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

			store[itemName] = store[itemName] - cnt
		}

		//fmt.Printf("params: %#v\n", params)
	}

	fmt.Println(store)
}
