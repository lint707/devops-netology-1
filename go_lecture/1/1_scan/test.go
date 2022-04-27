package main

import (
	"fmt"
	"math"
	"math/rand"
)

// var c, pyt bool
var i, j int = 2, 3

const Pi = 3.14

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 3 / 8
	y = sum - x
	return
}
func main() {
	// var i int
	var c, pyt, java = true, false, "no!"
	var x, y int = 7, 4
	//x, y := 7, 4

	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	k := 33
	var v int = 44 //change me!

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g prolems. \n", math.Sqrt(49))
	fmt.Println(math.Pi)
	fmt.Println(add(12, 8))
	fmt.Println(swap("text1", "new text"))
	a, b := swap("_txt_", "_not_")
	fmt.Println(a, b)
	fmt.Println(split(50))
	fmt.Println(i, c, pyt, j, java, k)
	fmt.Println(x, y, z, f)
	fmt.Printf("v is of typr %T\n", v)
	fmt.Println("v is of typr \n", Pi)
}
