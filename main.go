package main

import (
	"fmt"
	"./data"
	// "./reports"
	"./host"
)

func main() {
	fmt.Println()
	vector := data.GetVector(data.Dataa())
	// reports.GetComplete(vector)
	host.MainVector = vector
	// fmt.Println(vector)
	// host.Request()	
}
