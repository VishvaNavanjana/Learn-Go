package main

import "fmt"

func updateName(x *string){
	*x = "wedge"
}


func main(){
	name := "tifa"

	// fmt.Println("memory address of name is: ", &name)

	m := &name

	// updateName(m)

	fmt.Println(m)
	fmt.Println(*m)

	updateName(m)

	fmt.Println(name)
}