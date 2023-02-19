package main

import "fmt"

func main(){
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday","Monday", "Tuesday", "Wednesday","Saturday"}

	fmt.Println(days)

	//loops
	for d:=0; d < len(days); d++{
		fmt.Println(days[d])
	}

	for i := range days{
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v", index, day)
	}

	rougueValue := 1 

	for rougueValue < 10 {
		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	

}