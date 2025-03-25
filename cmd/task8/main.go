package main

import "fmt"

func main() {
	ch := make(chan bool, 2)
	waitRead := make(chan struct{})
	ch <- true
	go func() {
		for value := range ch {
			fmt.Println(value)
		}
		close(waitRead)
	}()
	ch <- true
	close(ch)
	<-waitRead
}
