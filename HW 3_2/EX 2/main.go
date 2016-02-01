package main

import "fmt"

func main(){
	funcexp := func(num int)(){
		var div2 int
		div2 = num/2
		if num%2 == 0{
			fmt.Println(div2,true)
		} else {
			fmt.Println(div2,false)
		}
		return
	}

	funcexp(1)
	funcexp(2)
}
