package main

import "fmt"

func main(){

	menu := map[string] float64{
		"soup" : 4.99,
		"pie" : 7.99,
		"salad": 6.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k,v := range menu {
		fmt.Println(k,"-",v)
	}

	//ints as key type
	phonebook := map[int] string{
		24353535 : "vishva",
		57463427 : "sasiri",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[57463427])

	phonebook[57463427] = "udith"

	fmt.Println(phonebook)

}