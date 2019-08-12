package main
//copy the slice and map
func main() {

	a := []int{10,20,30}
	b := make([]int,10)
	count :=copy(b, a)
	println("the number of elements copied " ,count)  //3
	println("the first element of new slice  " ,b[0])

	var c []int  //only declared
	countCopy :=copy(c, a)
	println("the number of elements copied " ,countCopy)  //0
	//so we cannot copy when it is not initialized

	//change the data of the b[0 ] to 100
	b[0]=100
	println("the first element of a slice after change to b " ,a[0]) // 10
	println("the first element of b slice after change to b  " ,b[0]) //100
	//so copy also create new one and not the same reference copy


	// here is just assignment
	map1 := map[string]bool{"A": true, "B": true}
	map2  := make(map[string]bool)

	map2 = map1

	println(len(map2))

	map2["B"]= false
	println("the first element of a map1 after change to map2 " ,map1["B"]) // false
	println("the first element of a map2 after change to map2 " ,map2["B"]) // false

	// here it is just an assignment so reference is copied







}
