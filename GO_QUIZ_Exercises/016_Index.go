package main

import "fmt"

func main() {
	// array but compiler count for you
	a := [...]string{
		"g",
		"o",
		7:"k",
		6:"u",
		4:"n",
		5:"g",
		3:"a",
		2:"l",
	} // ... means array but compilers count for you
	// print the type and value
	fmt.Printf("%T %[1]v",a) //[8]string [g o l a n g u k]
	b :=[] string{

		0:"hello",
		"test",
		3: "hey",
		2 :"gowtham",

	}// if we dont give the index , it will calculate the index based on where it located and not by other index
	//here test is located in 1, so it index is 1
	// if you have any other data with same index like 1:"hell"  it throws error saying duplicate index
	fmt.Println(b)


	
}
