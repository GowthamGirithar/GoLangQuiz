package main

import "flag"

//flag demo
// steps to execute :
//
//  > go run 010_Flag.go              // default value is considered
func main() {
	learningFlag := flag.Bool("isLearning" , true , "It is used to say it is for learning")
	var description string
	flag.StringVar(&description, "description", "Learning of flags ", "description about program")
	flag.Parse()  //this is required

	//>go run 010_Flag.go -isLearning=false --description=Learning_Flag //false and Learning_Flag
	//> go run 010_Flag.go -isLearning=false   //output is false and Learning of flags
	println("is learning flag " , *learningFlag)  // it is always address so we have to take the value
	println("is description flag " , description)  //here we can directly use becuase address is bound to the flag varaible
}
