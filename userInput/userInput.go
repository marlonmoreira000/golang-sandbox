package main

import "fmt"

func main() {
	fmt.Println("Enter player names (separated by space):")
	var p1, p2 string
	fmt.Scan(&p1, &p2)
	fmt.Println("Players: ", p1, p2)
}