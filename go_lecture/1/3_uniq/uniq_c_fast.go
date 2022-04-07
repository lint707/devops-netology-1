package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var counter int
	var prev string
	for in.Scan() {
		txt := in.Text()
		counter++
		if txt == prev {
			continue
		}
		if txt < prev {
			log.Fatal("file not sorted")
		}
		fmt.Printf("%d %s\n", counter, txt)
		counter = 0
		prev = txt
	}
}
