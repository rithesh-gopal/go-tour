package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	rand.Seed(89)
	fmt.Println("My favourite number is ", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}
