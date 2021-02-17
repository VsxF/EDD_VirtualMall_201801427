package main

import (
	"fmt"
	"./data"
	// "./reports"
	"./host"
)

func main() {
	fmt.Println()
	vector := data.NewVector()
	vector.GetVector(data.Dataa())
	fmt.Println(vector)
	// reports.GetComplete(vector)
	host.MainVector = vector
	
	host.Request()	
}
