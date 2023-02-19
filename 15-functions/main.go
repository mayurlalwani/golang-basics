package main

import "fmt"

func main(){
	fmt.Println("Functions in Golang")
	greeter()
	result := adder(3,5)
	fmt.Println("Result is:", result)
	proResult := proAdder(3,5,6,4,7,8,9)
	fmt.Println(proResult)

}

func adder(a int,b int) int {
	return a + b
}

func proAdder(values ...int) int{
	total := 0
	for _, value :=range values {
		total+=value
	}
	return total
}

func greeter(){
	fmt.Println("Hello from Golang")
}