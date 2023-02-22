package main

import "fmt"

func main() {
	a, b := twoReturn(2, 5)
	fmt.Println(a, b)
}

func twoReturn(one int, two int) (int, int) {
	return one + two, one - two
}