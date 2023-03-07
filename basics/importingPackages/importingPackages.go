package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	n1, err1 := strconv.Atoi(os.Args[1])
	n2, err2 := strconv.Atoi(os.Args[2])
	fmt.Printf("%v + %v = %v\n", n1, n2, add(n1, n2))
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
}
