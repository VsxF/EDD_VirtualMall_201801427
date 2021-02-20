package main

import (
	"fmt"
	"./data"
	"./host"
)

func main() {
	fmt.Println()
	vector := data.NewVector()
	data, _ := data.Dataa()
	
	vector.GetVector(data)
	host.MainVector = vector
	
	host.Request()	
}