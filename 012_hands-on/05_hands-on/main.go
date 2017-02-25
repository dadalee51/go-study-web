package main

import (
	"log"
	"os"
	"html/template"
)

 

type menuitem struct{
	Type string
	MenuItem string
	Price float64
}

type menus [] menuitem

type restaurant struct{
	Menus menus
	Address string
}

var tpl *template.Template

var Restaurants [] restaurant
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}



func main(){
	
	Restaurants = []restaurant{
		restaurant{
			Menus: []menuitem{
				menuitem{
					Type: "Breakfast",
					MenuItem: "Hashbrown",
					Price : 6.00,
				},

				menuitem{
					Type: "Lunch",
					MenuItem: "Burger",
					Price: 10.00,
				},
			},
			Address: "99 sunday Street",
		},
	}

	err := tpl.Execute(os.Stdout, Restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}