package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func main() {
	var wait sync.WaitGroup
	counter := 200000
	wait.Add(counter)
	// go hello()
	// go hello()
	// wait.Wait()
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
	defer wait.Wait()

}
func hello(wait *sync.WaitGroup, counter int){
	fmt.Println(counter)
	wait.Done()
}