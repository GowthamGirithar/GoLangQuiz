package main

import "log"

func main() {
	log.Fatal("error")  //os.Exit(1) - so application is terminated - abnormal termination
	println("The error is printed")
}


//output is only error and not the next line