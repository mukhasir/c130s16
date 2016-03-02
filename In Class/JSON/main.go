package main

import (
	"encoding/json"
	"fmt"
)

/*type Person struct{
	Name string
	Age int
}*/
type candyStore struct {
	Name    string
	Candies map[string]int
}

func main() {
	/*p := Person{"Mukhasir",25}
	bs, err := json.Marshal(p)
	if err !=nil{
		panic(err)
	}
	data := `Person{"Name":"Mukhasir","Age":25}`
	var p Person
	err := json.Unmarshal([]byte(data),p)
	if err !=nil{
		panic(err)
	}
	js:=string(data)
	fmt.Println(js)*/
	st2 := candyStore{
		Name:    "Circle M",
		Candies: map[string]int{"mnms": 3, "snickers": 7},
	}
	fmt.Println(st2)
	bs, err := json.Marshal(st2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	var st1 map[string]interface{}
	json.Unmarshal(bs,&st1)
	fmt.Println(st1["Candies"].(map[string]interface{})["mnms"])
	var st3 candyStore
	json.Unmarshal(bs,&st3)
	fmt.Println(st3.Name)

}
