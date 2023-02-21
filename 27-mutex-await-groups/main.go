package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Race Condition")
	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	mut.RLock()

	wg.Add(3)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("One routine")
		mut.Lock()
		score = append(score,1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Two routine")
		mut.Lock()
		score = append(score,2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("Three routine")
		mut.Lock()
		score = append(score,3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}