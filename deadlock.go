package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		for _, value := range array {
			ch <- value
		}
		close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println(value)
			wg.Done()
		}
	}()

	wg.Wait()
}
