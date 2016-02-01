package main

import "fmt"

func main(){
	fmt.Println(findMax(1,2,3,4,78,34,21,67))
}

func findMax(nums... int) (int) {
	var max int
	for _,num := range nums{
		if num > max{
			max = num
		}
	}
	return max
}
