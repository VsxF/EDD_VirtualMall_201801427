package main

import (
	data "./data/stores"
	"./data/orders"
	// "./reports"
	"./host"
)

func main() {
	vector := data.NewVector()
	data, _ := data.Dataa()
	
	vector.GetVector(data)

	vector.AuxSetInventrorys()
	host.MainVector = vector

	calendar := orders.NewCalendar()
	calendar.AuxSetOrders()
	host.MainOrders = calendar

	// // reports.GetDayOrders(calendar, "13-08-2020")
	// reports.GetMatrixGraph(calendar, "13-11-2018")

	host.Request()	
}