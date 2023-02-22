package main

import "fmt"

func main() {
	for i, v := range []int{1,2,3,4,5,6,7,8,9} {
		if i % 2 == 0 {
			continue
		}
		if i == 7 {
		break
		}
		fmt.Println(i, v)
	}
}