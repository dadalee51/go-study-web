package main

import (
	"log"
	"os"
	"html/template"
	"fmt"
	"rob"
)

type hotel struct{
	Name, Address, City, Zip string
}

type region struct{
	Region string
	Hotels []hotel
}


var tpl *template.Template
var Regions []region

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}



func main(){
	fmt.Println("main - 01")
	rob.Demo()

	Regions = []region{
		region{
			Region: "Southern",
			Hotels : [] hotel{
				hotel{
					Name: "BedVilla",
					Address:"1123 Gen Street",
					City: "Vegsbane",
					Zip: "90210",
				},
				hotel{
					Name: "BedVilla2",
					Address:"1124 Gen Street",
					City: "Meatsbane",
					Zip: "90211",
				},
			},
		},

		region{
			Region: "Central",
			Hotels : [] hotel{
				hotel{
					Name: "ConVilla",
					Address:"3323 Gen Street",
					City: "Fruitsbane",
					Zip: "90220",
				},
				hotel{
					Name: "Minilla2",
					Address:"333 seet street",
					City: "Medacbane",
					Zip: "34411",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, Regions)
	if err != nil {
		log.Fatalln(err)
	}
}