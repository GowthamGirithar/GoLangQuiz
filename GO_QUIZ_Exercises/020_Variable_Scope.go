package main

import "strconv"

func main() {

	i := 1
	s :="1000"
	if len(s) >1 {
		i , _ := strconv.Atoi(s)
		i =i+5
	}
	println(i) // output will be 1 , because i inside the if block is redefined and that have the scope only inside the if block
	// i , _ := strconv.Atoi(s) to be changed to i , _ = strconv.Atoi(s)

	
}
