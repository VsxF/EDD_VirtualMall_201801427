package orders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type AuxOrders struct {
	Orders []Order `json:"Pedidos"`
}


/// JUST FOR Develop
func (calendar *Calendar) AuxSetOrders() {
	jsonFile, _ := os.Open("./Pedidos.json")
	defer jsonFile.Close()

	var ordersJSON AuxOrders

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err2 := json.Unmarshal([]byte(byteValue), &ordersJSON)

	if err2 != nil {
		fmt.Println("error: ", err2)
	}
	calendar.SetOrders(ordersJSON)
	//fmt.Println(calendar)
}

func (calendar *Calendar) SetOrders(AuxOrders AuxOrders) {
	for i := 0; i < len(AuxOrders.Orders); i++ {
		year, months := mapYear(AuxOrders, i)
		
		calendar.InsertYear(year, months)
	}
}

func mapYear(AuxOrders AuxOrders, k int) (int, *Months) {
	months := NewMonthsList()

	monthDate, year := getMonth(AuxOrders.Orders[k].Date)
	orders := getOrders(AuxOrders, k)

	months.setMonth(monthDate, orders)
	
	return year, months
}

func getOrders(AuxOrders AuxOrders, k int) *Orders {
	orders := NewOrders()
	aux := AuxOrders.Orders[k]
	InsertarOrden(orders, aux.Date, aux.Store, aux.Department, aux.Qualification, aux.Products)
	
	return orders
}

//return Months.Date, year
func getMonth(date string) (int, int) {
	aux := []byte(date)
	
	yearSTR := string(aux[6]) + string(aux[7]) + string(aux[8]) + string(aux[9])
	monthSTR := string(aux[3]) + string(aux[4])
	
	month, _ := strconv.Atoi(monthSTR)
	year, _ := strconv.Atoi(yearSTR)

	return month, year
}