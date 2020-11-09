package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{5, 10}
	fmt.Println(v.X)

	p := &v
	fmt.Println(p.Y)
}
