package main

import (
	"./data"
	"./host"
)

func main() {
	vector := data.NewVector()
	data, _ := data.Dataa()
	
	vector.GetVector(data)
	host.MainVector = vector
	
	host.Request()	
}