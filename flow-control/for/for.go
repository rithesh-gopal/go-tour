package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	runs := 1
	for runs < 100 {
		runs += runs
	}
	fmt.Println(runs)
}
