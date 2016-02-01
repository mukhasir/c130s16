package main

import "fmt"

func main() {
	var smallNum int
	var bigNum int
	fmt.Println("Enter Small Number")
	fmt.Scan(&smallNum)
	fmt.Println("Enter Big Number")
	fmt.Scan(&bigNum)
	fmt.Println(bigNum % smallNum)
}
