package main
//swap number in golang
func main() {
	var a , b=10,20
	println("the data before  swap ", a , b)

	//swap the data
	a, b = b, a
	println("the data after swap " ,a , b) //20 , 10
}
