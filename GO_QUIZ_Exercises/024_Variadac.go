package main

func GetData(data []int)  {
	data[0]=100
}
// normally variadac function consider data as slice
func GetData1(data... int)  {
	data[0]=10
}


func main() {
	i := []int{10,20,30}
	GetData(i)
	println(i[0]) //100 , becuase normally slice is pass by reference

	GetData1(i...)// normally variadac fun creates new slice , since we have maintained i ... (three dots), it dont create the new slice
	println(i[0]) //10


}
