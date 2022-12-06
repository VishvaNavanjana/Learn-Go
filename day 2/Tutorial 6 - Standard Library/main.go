package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "hello there!"

	// fmt.Println(strings.Contains(greetings,"hello"))
	// fmt.Println(strings.ReplaceAll(greetings,"hello","hi"))
	// fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Split(greetings, " "))

	// fmt.Println(strings.Index(greetings,"ll"))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)

	fmt.Println(index)

	name := []string{"yoshi", "mario", "peach", "bowser"}
	sort.Strings(name)
	fmt.Println(name)

	fmt.Println(sort.SearchStrings(name, "bowser"))
}
