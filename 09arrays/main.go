package main

import "fmt"

func main(){
	fmt.Println("Arrays in golang")

	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "Chickoo"
	fruitList[4] = "Watermelon"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list is:", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list is:",vegList)

}