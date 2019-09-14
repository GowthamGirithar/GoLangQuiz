package main

func PrintData(i int)  {
	println(i)
}

func main() {
	i := 10
	defer PrintData(i)
	i= i+10
	
}
// output will be 10 , because defer will function will be evaluated at the time of execution itself and not at the time of actual call