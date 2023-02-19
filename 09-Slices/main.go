package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Slices in GoLang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int,4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 400
	highScores[3] = 589

	fmt.Println(highScores)

	highScores = append(highScores, 300, 500, 789)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}