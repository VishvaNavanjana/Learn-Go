package main

import "fmt"

func main(){
	//arrays cant change length
	// var ages [3] int = [3]int{20,25,35}
	var ages = [3]int{20,25,35}

	name := [4]string{"yoshi","mario","peach","bowser"}
	name[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(name,len(name))

	// slices (use arrays under the hood)
	var scores = []int{100,50,60}
	scores[2] = 25

	scores = append(scores,85)

	fmt.Println(scores,len(scores))

	// slice ranges
	rangeOne := name[1:3]
	rangetwo := name[2:]
	rangeThree := name[:3]

	fmt.Println(rangeOne,rangetwo,rangeThree)
}