package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name		string `json:"courseName"`
	Price		int
	Platform	string	`json:"website"`
	Password	string	`json:"-"`
	Tags		[]string `json:"tags,omitempty"`
}

func main(){
	fmt.Println("Json in Golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "Learncodeonline.in", "Test@123", []string{"web-dev","js"}},
		{"MERN Bootcamp", 199, "Learncodeonline.in", "Test@789", []string{"full-stack","js"}},
		{"Golang Bootcamp", 499, "Learncodeonline.in", "Test@456", []string{"backend","go"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s",finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"courseName": "Golang Bootcamp",
		"Price": 499,
		"website": "Learncodeonline.in",
		"tags": [
				"backend",
				"go"
		]
	}
	`)

	var lcoCourse course 

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else{
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v\n", k,v)
	}


}