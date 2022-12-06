package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string){
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string){
	fmt.Printf("Goodbye %v \n", n)
}

func cycleName( n []string, f func(string)){
	for _,v := range n {
		f(v)
	}
}


func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {
	// sayGreetings("vishva")
	// sayGreetings("sasiri")
	// sayBye("udith")

	name := []string{"yoshi", "mario", "peach", "bowser","luigi"}

	cycleName(name,sayGreetings)

	a1 := circleArea(10.5)
	fmt.Printf("%0.3f",a1)
}