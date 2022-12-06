package main

import "fmt"

func main(){
	//strings
	var nameOne string = "vishva"
	var nametwo = "navanjana"
	var nameThree string

	fmt.Println(nameOne,nametwo,nameThree)

	nameThree = nameOne

	nameOne = "clouda"

	fmt.Println(nameOne,nametwo,nameThree)

	//without var
	namefour := "apple"

	fmt.Println(namefour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne,ageTwo,ageThree)

	//bits & memory
	// var numOne int8 = 127
	// var numTwo uint8 = 255

	var scoreOne float64 = 25.98 //most time
	scoreTwo := 1.54 

	fmt.Println(scoreOne,scoreTwo)


}