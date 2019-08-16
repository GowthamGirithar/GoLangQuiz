package main

import (
	"net/http"
	"sync"
)
//this data is maintained throughout the application and not only at the time of request
//it is like static context
var StaticData string
var doOnce sync.Once


func main() {
	http.HandleFunc("/test" , testFn)
	http.ListenAndServe(":8080" , nil)
}

func testFn(writer http.ResponseWriter, request *http.Request) {
	doOnce.Do(func(){
		StaticData="Apple"
		println("inside once block")
	})
	println("inside test function" , StaticData)
	writer.Write([]byte("Hello enjoy learning Sync do once"))

}

/**

so only at the first request it executed aftewr that any number of request it didnt do
inside once block
inside test function
inside test function
inside test function
inside test function
inside test function
inside test function
 */