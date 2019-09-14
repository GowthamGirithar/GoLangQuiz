package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]  // a[low:high:max]
	fmt.Println(t[0]) // 4
	t1 :=a[2:3:5] // 5 is the capacity
	fmt.Printf("%v" , t1)
	fmt.Println(cap(t1))
}
