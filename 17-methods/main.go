package main

import "fmt"

func main(){
	fmt.Println("Methods in Golang")

	john := User{"John", "john@gmail.com", true, 28}
	fmt.Println(john)
	fmt.Printf("John's details are: %+v\n", john)
	john.GetStatus()
	john.NewMail()
}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}