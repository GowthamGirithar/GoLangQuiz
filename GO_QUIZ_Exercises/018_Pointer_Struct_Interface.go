
package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func h(x *S) {
}


//So you're confusing two concepts here. A pointer to a struct and a pointer to an interface are not the same.
//TODO: LOOKS BOTH SAME , YET TO VERIFY
func main() {
	s :=S{}
	p := &s
	f(s) //A
	g(&s) //B -> compile error
	f(p) //C
	g(p) //D -> compile error
	h(&s)
	h(s)//compile time error
	
}
