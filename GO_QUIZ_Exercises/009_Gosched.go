package main

import (
	"fmt"
	"runtime"
)

func main() {
	println(runtime.GOMAXPROCS(-1)) //4 in my system
	runtime.GOMAXPROCS(1) //changing to 1 - so only one thread there and we cannot switch context unless call Gosched
	go say("world")
	say("hello")
}

//As I said, when GOMAXPROCS variable is not specified, Go runtime is only allowed to use one thread,
// so it is impossible to switch execution contexts while goroutine is performing some conventional work,
// like computations or even IO (which is mapped to plain C functions). The context can be switched only
// when Go concurrency primitives are used, e.g. when you switch on several chans, or (this is your case) when you explicitly tell the scheduler to switch the contexts - this is what runtime.Gosched is for.
//when you explicitly tell the scheduler to switch the contexts - this is what runtime.Gosched is for.
func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		runtime.Gosched() // it will explicity say change the context
	}
}



