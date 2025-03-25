package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	count := 3
	wg.Add(count)
	for i := range count {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for v := range ch {
		sum += v
	}

	fmt.Printf("result: %d\n", sum)

}
