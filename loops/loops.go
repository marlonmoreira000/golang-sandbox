package main

import "fmt"

func main() {
	// type 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// type 2
	var i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// type 3
	var arr = []string{"one", "two", "three"}
	for i, v := range arr {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}
}