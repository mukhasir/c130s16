package main

import "fmt"

func main() {
	GetValue(9)
	fmt.Println("Yes")
}
func GetValue(num int) {
	var buffer []int
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			buffer[i-1] = i
		}
	}
	fmt.Println(buffer)
	getSpecVal(buffer)
	return
}
func getSpecVal(buffer []int) {
	fmt.Println(len(buffer))
}
