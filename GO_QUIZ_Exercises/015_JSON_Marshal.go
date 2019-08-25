package main

import "encoding/json"

type Sample struct {
	X int `json:"Key_1"`
	Y int `json:"Key_2,OmitEmpty"`
	Z int `json:"-"`
}

func main() {
	s:= Sample{
		X: 1,
		Y: 0,
		Z: 0,
	}
	data , err :=json.Marshal(s)
	if err != nil{
		println(err.Error())
	}
	println(string(data))
	
	
}

/**
Output
{"Key_1":1,"Key_2":0}

omit the part after , in json
omit the data itself if it is -
 */
