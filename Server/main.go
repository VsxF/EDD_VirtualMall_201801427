package main

import (
	data "./data/stores"
	"./host"
)

func main() {
	vector := data.NewVector()
	data, _ := data.Dataa()
	
	vector.GetVector(data)

	vector.AuxSetInventrorys()
	host.MainVector = vector

	host.Request()	
}