package main

import "fmt"

func main(){
	age := 25
	name := "adam"
	fmt.Println("my name is",name,"and my age is ",age)

	// Printf (formated strings)
	fmt.Printf("my age is %v and my name is %v \n",age,name)
	fmt.Printf("my age is %q and my name is %q \n",age,name)
	fmt.Printf("age is of type %T \n",age)
	fmt.Printf("you scored %0.2f points \n",225.5534)

	// Sprintf (save formated strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n",age,name)

	fmt.Println("the saved string is:",str)
}