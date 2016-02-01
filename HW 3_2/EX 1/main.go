package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(number int) (div2 int, even bool) {
	div2 = number / 2
	if number%2 == 0 {
		even = true
	} else {
		even = false
	}
	return div2,even
}
