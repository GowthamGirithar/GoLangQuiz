package main

func hello() []string{
	return nil
}

func main() {
	f := hello
	if f == nil{
		println("nil")
	}else{
		println("not nil") // not nil will be printed , because function is only defined and it has memory addesss, if it is hello() then ans will be nil
	}
	
}
