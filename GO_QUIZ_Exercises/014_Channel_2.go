package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	ch :=make(chan int, 3)

	go TestChannel(ch)
	println("the code after routine")
	for i := range ch{  // RANGE IN CHANNEL IS BLOCKING
		println(i)
	}
	
}

func TestChannel(ch chan int) {
	for i:=0;i<3;i++{
		ch <- i
		println("the channel value is " , i)
	}
	defer close(ch)
}

/**
read from channle block , so output is like below

the code after routine
the channel value is  0
the channel value is  1
the channel value is  2
0
1
2
 */
