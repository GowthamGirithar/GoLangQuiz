package main

type Employee struct{
	Name string
}

//comparing the two struct values
func main() {
	emp1 := Employee{}
	emp1.Name="Gowtham"

	emp2 := Employee{}
	emp2.Name="Gowtham"

	println("the emp1 and emp2 are equal ?" , emp1 == emp2)
}
