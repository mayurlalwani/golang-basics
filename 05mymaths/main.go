package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main(){
	fmt.Println("Welcome to maths in Golang")

	// var mynumberOne int = 2
	// var munumberTwo float64 = 4.5
	// fmt.Println("The sum is: ", mynumberOne + mynumberTwo)

	//How to generate random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	//Random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)



}