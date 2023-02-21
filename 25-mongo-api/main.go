package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mayurlalwani/mongoapi/router"
)

func main(){
	fmt.Println("MongoDb API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	http.ListenAndServe(":4000", r)
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at PORT 4000...")

}