package main

import "fmt"

func main() {
	arr := [3]string{"a", "b", "c"}
	sl1 := arr[1:]
	sl1 = append(sl1, "z")
	fmt.Println(sl1)

	var numbers []int
	fmt.Printf("%T", numbers)

	sl2 := make([]string, 3)
	copy(sl2, sl1)
	fmt.Println(sl2)
}