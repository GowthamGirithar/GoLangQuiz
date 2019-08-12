package main

// alias
import "math/rand"
import cryptRand "crypto/rand"
// . to use the method without package name in code
import . "fmt"
// we are not using and only for side effects
import _ "github.com/go-sql-driver/mysql-master"
func main() {


	//import with alias name so no need to use the same name
	//it will be used when you have two package name with same name to resolve what we are using
	println(rand.Int())
	//this rand is used with alias name
	randomBytes := make([]byte, 32)
	_, err:= cryptRand.Read(randomBytes)
	if err != nil{

	}

	// if you want to use the methods everytime without package name in code we can use . in the import
	Printf("the method is used without the package name")


	// _ in import is used if it is not used in code by us but internally it requires
	//example sql driver . go code internally requires for loading the drivers we dont want to anything



}
