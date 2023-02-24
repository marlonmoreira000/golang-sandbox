package main

import "fmt"

func errorHandler() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered ", r)
	}
}

func Divide(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("can't divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	no := Divide(10, 0)
	fmt.Println(no)
	no = Divide(10, 1)
	fmt.Println(no)
}
