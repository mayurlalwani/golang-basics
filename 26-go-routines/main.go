package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main(){
	// go greeter("Hello")
	// greeter("world")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.in",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string){
	defer wg.Done()
	res, err := http.Get(endpoint)
	
	if err != nil{
		fmt.Println("Incorrect endpoint")
	}
	mut.Lock()
	signals = append(signals,endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	
}

// func greeter(greet string){
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(greet)
// 	}
// }