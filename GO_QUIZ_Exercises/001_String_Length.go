package main

import "unicode/utf8"

//len function is used to get the length by bytes
//for special character we could see the differences, so Rune character in string is used
func main() {

	var dataStr string ="文字"
	println(len(dataStr))   //6 due to 6 bytes

	for i:=0 ; i< len(dataStr) ; i++{
		//iterate 6 times
	}

	//CORRECT ONE
	//to avoid the above behaviour
	length := utf8.RuneCountInString(dataStr)
	println(length) //2 - it will count the character


	for i:=0 ; i< length ; i++{
		//iterate 2 times
	}
}
