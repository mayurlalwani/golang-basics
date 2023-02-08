package main

import "fmt"

func main(){
	fmt.Println("Structs in Golang")

	john := User{"John", "john@gmail.com", true, 28}
	fmt.Println(john)
	fmt.Printf("John's details are: %+v", john)
}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}