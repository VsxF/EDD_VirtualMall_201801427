package main

import (
	"fmt"
	"./data"
	"./reports"
)

func main() {
	fmt.Println("main")
	vector := data.VectorMain()
	//fmt.Println(vector.Vector)
	reports.GetComplete(vector)
	//reports.GetComplete()
}
