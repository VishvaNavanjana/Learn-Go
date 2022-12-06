package main

import "fmt"

func updateName(x string) string{
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["cofee"] = 2.99
}

func main(){
	name := "tifa"
	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64 {
		"pie":5.95,
		"ice cream":3.99,
	}

	updateMenu(menu)

	fmt.Println(menu)
}