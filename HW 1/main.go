package main

import (
	"html/template"
	"log"
	"os"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

func main() {
	var err error

	myName := Name{First: "Mukhasir", Middle: "Shah", Last: "Syed"}

	tmp_Test := template.New("temp_Cond_logic.html")

	tmp_Test, err = tmp_Test.ParseFiles("temp_Cond_logic.html")

	if err != nil {
		log.Fatalln(err)
	}

	err = tmp_Test.Execute(os.Stdout, myName)

	if err != nil {
		log.Fatalln(err)
	}
}
