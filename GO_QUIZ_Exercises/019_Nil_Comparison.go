
package main

func main() {
	var m map[string]int //A
	m["a"] = 1
	if v := m["b"]; v != nil { //B//compile type error , since int cannot be checked to nil
		println(v)
	}


	var m1 map[string]interface{} //c
	m1["a"] = 1
	if v := m1["b"]; v != nil { //D//but if we have the interfacetype , then it can be checked with nil
		println(v)
	}
}
//compile type error , since int cannot be checked to nil
//but if we have the interfacetype , then it can
