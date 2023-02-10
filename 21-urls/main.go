package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asdffsd"

func main(){
	fmt.Println("URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are:%T", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is: ", val)
	}

	partsofUrl := &url.URL{
		Scheme: "https",
		Host:	"lco.dev",
		Path: 	"/tutcss",
		RawPath: "user=mayur",
	}

	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}