package main

import "sort"

type EmployeeDetail struct{
	Name string
	age int
}

//sorting the slice
func main() {
	// sorting the slice of type string
	var dataStr =[]string{"Apple", "Orange", "Bat"}

	//to check whether it is already sorted or not
	println("the sorted or not ", sort.StringsAreSorted(dataStr))  //false
	sort.Strings(dataStr)

	for d := range dataStr{
		println(dataStr[d])
	}
	// sorting the slice of type int
	var dataInt =[]int{100, 10, 0}
	sort.Ints(dataInt)

	for d := range dataInt{
		println(dataInt[d])
	}

	//reverse sort

	sort.Sort(sort.Reverse(sort.IntSlice(dataInt)))

	// sorting the slice which is struct type

	emp :=[]EmployeeDetail{
		{"B", 24},{"ab", 24},
	}

	sort.Slice(emp, func(i, j int) bool {
		return emp[i].Name > emp[j].Name
	})

	for e := range emp{
		println(emp[e].Name)
	}





}
