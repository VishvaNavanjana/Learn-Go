package main

import "fmt"

func main(){
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is :", x)
	// 	x++
	// } 

	name := [4]string{"yoshi","mario","peach","bowser"}

	for i := 0; i < len(name); i++{
		fmt.Println(name[i])
	}

	for index, value := range name {
		fmt.Printf("the position at index %v is %v\n",index,value)
	}

	for _, value := range name {
		fmt.Printf("the value is %v\n",value)
	}

}