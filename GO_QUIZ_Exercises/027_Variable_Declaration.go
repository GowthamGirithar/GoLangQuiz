package main

func main() {

	var a int
	a, err := test()
	if err != nil {

	}
	// if any one of the variable is newly declared , then we have to use :=
	//otherwise =
	// if the case is a, _ = test() -> correct , nothing is new to be declared
	a, err1 := test()
	//a, err1 := test()  > will give compile time error
	if err1 != nil {

	}
	println(a)
	var b int
	b, _ = test()
	println(b)

}

func test() (int, error) {
	return 10, nil
}
