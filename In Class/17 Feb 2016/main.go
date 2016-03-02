package main

import (
	"fmt"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type hotels struct {
	Hotels []hotel
}

/*func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}*/

func main() {
	h := hotels{
		Hotels: []hotel{
			hotel{
				Name:    "First",
				Address: "Fresno",
				City:    "Fresno",
				Zip:     93726,
				Region:  "Central",
			},
			hotel{
				Name:    "Second",
				Address: "Fresno",
				City:    "Fresno",
				Zip:     93726,
				Region:  "Central",
			},
		},
	}
	fmt.Println(h)
}
