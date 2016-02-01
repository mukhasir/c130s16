package main

import (
	"fmt"
	"reflect"
)

func main(){
	var num int
	var num1 float32
	var val string
	valchar := 'A'
	var get bool
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(val))
	fmt.Println(reflect.TypeOf(valchar))
	fmt.Println(reflect.TypeOf(get))
}
