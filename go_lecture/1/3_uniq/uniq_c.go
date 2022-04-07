package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	counters := map[string]int{
		"test": 0,
	}
	for in.Scan() {
		txt := in.Text()
		counters[txt]++
	}
	values := make([]string, 0, 100)
	for val, _ := range counters {
		values = append(values, val)
	}
	// sort.Strings(values)
	sort.Slice(values, func(i, j int) bool { // это для сортировки по номеру
		return counters[values[i]] < counters[values[j]]
	})
	fmt.Println(counters, values)
	var val string
	for i := 0; i < len(values); i++ {
		val = values[i]
		fmt.Printf("%d %s\n", counters[val], val)
	}
}
