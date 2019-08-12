package main

import (
	"unsafe"
)

type EmptyStruct struct{
	Name string
}

type S struct {
	A struct{}
	B struct{}
}

//why we would prefer the empty struct over the others
func main() {
	data := struct{}{}  // eg: Employee{} -> struct{}{}
	println(unsafe.Sizeof(data))  // 0

	// here we just declared like above but the memory is occupied
	var d EmptyStruct
	println(unsafe.Sizeof(d))  //16

	// string it occupies 16 so above returned by that memory
	var str string
	println(unsafe.Sizeof(str))  //16

	// struct with data members struct
	var s S
	println(unsafe.Sizeof(s))  //0


	// if you are traversing the data and want to set as read , then use the empty struct
	//signalling the channel

	mapData := make(map[int]struct{})
	for i:=0 ; i <10; i++{
		if i%2==0{
			mapData[i] = struct{}{}
		}
	}

	for _,value := range mapData{
		if value == struct{}{}{
			println("the data is empty struct")
		}

	}



}
