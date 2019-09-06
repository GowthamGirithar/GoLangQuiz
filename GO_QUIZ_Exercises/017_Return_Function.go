package main

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func main() {
	var s1=S{}
	print(s1) // empty struct cannot print
	//print(InitType())
	println(InitPointer()) //returns address so nil
	println(InitEfaceType() ) // return function , it will have address
	println(InitEfacePointer())// return function , it will have address
	println(InitIfaceType() )// return function , it will have address
	println(InitIfacePointer() )// return function , it will have address


	println(InitPointer() == nil) //true
	println(InitEfaceType() == nil)
	println(InitEfacePointer() == nil)
	println(InitIfaceType() == nil)
	println(InitIfacePointer() == nil)
}


