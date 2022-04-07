package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// запускать cat data.txt | go run grep.go shell
// запускать cat data.txt | go run grep.go -v shell

func main() {
	// fmt.Println(os.Args)
	substr := os.Args[len(os.Args)-1]
	isReverse := false
	for i, val := range os.Args {
		if val == "-v" {
			isReverse = true
			fmt.Println("-v found on pos", i)
		}
	}
	fmt.Println(substr, os.Args)
	in := bufio.NewScanner(os.Stdin)
	// END:
	for in.Scan() {
		txt := in.Text()
		txt = strings.ToLower(txt)
		isFound := strings.Contains(txt, substr)
		// fmt.Println(txt, isFound)
		if (isReverse && !isFound) || (!isReverse && isFound) {
			fmt.Println(txt)
		}
	}
}
