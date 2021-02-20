package main

import (
	"fmt"
	"./data"
	"./reports"
	"./host"
)

func main() {
	fmt.Println()
	vector := data.NewVector()
	data, _ := data.Dataa()
	
	vector.GetVector(data)
	reports.SaveVector(*vector)
	host.MainVector = vector
	
	host.Request()	
}