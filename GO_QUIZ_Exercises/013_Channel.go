package main

import (
	"fmt"
)

func main() {
	ch :=make(chan int, 1)
	go func() {
		ch <- 1
		fmt.Println("Running goroutine")
	}()
	<- ch
	fmt.Println("End Main")
	
}
/**
output will be
Running goroutine
End Main

even in buffered channel , reading from channel is blocking
 */
