package main

import (
	"fmt"
	// "tidwall/sjson"
	// "reflect"
	"encoding/json"
	"os"
	// "example.com/proj/go-proj/helper"
)


var Val int

type Analysis struct{
	Name string `json: "name"`
	Values int `json : vals`
	Age int `json : "age"`
}

type NewData struct {
	Name string `json:"name"`
	Data  int  `json:"data"`
	Val [] int `json:val`
}
	




func main()  {

	js, _ := os.ReadFile("data.json")
	var newdata NewData
	data := json.Valid(js)
	fmt.Println(data)


	if data{
		_ = json.Unmarshal([]byte(js), &newdata)
		
		fmt.Println(newdata.Val)

		
	}
}



