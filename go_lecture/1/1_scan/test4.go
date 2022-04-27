package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
	s         string
}

var m map[string]Vertex
var m1 map[int]Vertex
var m2 = map[string]Vertex{
	"bella": { //Vertex{
		12, 34, "rr",
	},
	"googla": { //Vertex{
		56, 32, "qq",
	},
}

func main() {
	fmt.Println("test")
	m = make(map[string]Vertex)
	m["Bell"] = Vertex{
		32, 24, "test",
	}
	fmt.Println(m["Bell"])
	m1 = make(map[int]Vertex)
	m1[12] = Vertex{
		11, 22, "text2",
	}
	fmt.Println(m1[12])
	fmt.Println(m2)

	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	v1, ok1 := m3["Answer"]
	fmt.Println("The test value:", v1, "Present?", ok1)

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
