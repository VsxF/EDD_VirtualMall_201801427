package main

import (
	"fmt"
	"./data"
	// // "./reports"
	"./host"
)

func main() {
	fmt.Println()
	vector := data.NewVector()
	data, _ := data.Dataa()
	// fmt.Println(a)
	
	vector.GetVector(data)
	// fmt.Println(vector)
	// reports.GetComplete(vector)
	// reports.SaveVector(*vector)
	host.MainVector = vector
	
	host.Request()	
}