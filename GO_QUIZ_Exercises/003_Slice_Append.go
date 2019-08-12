package main

func main() {
	//normally slice is the reference one and if we copy it will copy the reference
	var data =[]int{10,20,30}
	newData := data[:]   // nweData will have now 20 and 30
	println(" the data after copying ", data[0])
	println(" the newdata after copying ",newData[0])


	//change the value of new data index 0
	newData[0]=30

	//since the slice is referece copy it will change the value to old one also
	println(" the data after chanage ", data[0])
	println(" the newdata after chanage ",newData[0])


	// append use
	appendData := append(data, 20)
	println(" the appendData initial value ", appendData[0])
	appendData[0] =50

	println(" the appendData after change value ", appendData[0])
	println(" the data after change value in append data", data[0]) // still print 10

	// print the address
	//each slice location stores the address
	println(" address of data ", &data[0])
	println(" address of data ", &newData[0])
	println(" address of data ", &appendData[0])
	//so final conclusion append will create the new thing and not reference
}
