package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex1 struct {
	X int
	Y int
	Z float32
	O string
}
type Vertex struct {
	X int
	Y int
}

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Printf("test text")
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j, i)
	fmt.Println(Vertex1{1, 3, math.Pi, "TEXT"})
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 7
	q := &v
	q.X = 1e4
	fmt.Println(v.X, Vertex{v.X, v.Y}, v)
	var a [2]string
	a[0] = "txt"
	a[1] = "second"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	prim := [6]int{2, 4, 6, 1, 5, 9}
	fmt.Println(prim)
	var s []int = prim[2:5]
	fmt.Println(s)
	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var z []int
	printSlice(z)

	z = append(z, 0)
	printSlice(z)

	z = append(z, 2, 4, 5)
	printSlice(z)

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 4)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		fmt.Println(i, pow[i])
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(z []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(z), cap(z), z)

}
